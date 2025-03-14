package weather

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"weather-api/internal/config"
	"weather-api/internal/weather/model"

	"github.com/google/uuid"
	pq "github.com/lib/pq"
)

type Repository interface {
	GetWeatherByLocation(ctx context.Context, location string) (float32, error)
}

type repository struct {
	db      *sql.DB
	connStr string
}

func NewRepository(configPostgres config.Postgres) Repository {
	var connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		configPostgres.Host, configPostgres.Port, configPostgres.User, configPostgres.Password, configPostgres.Name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-shutdownChan
		db.Close()
		os.Exit(0)
	}()

	err = db.Ping()
	if err != nil {
		log.Fatal("Could not connect to the db", err)
	}

	log.Println("Successfully connected to the db")

	return &repository{db: db, connStr: connStr}
}

func (r *repository) GetWeatherByLocation(ctx context.Context, location string) (float32, error) {
	return r.getWeatherByLocationFromDB(ctx, location)
}

var (
	mutex sync.Mutex
)

func (r *repository) getWeatherByLocationFromDB(ctx context.Context, location string) (float32, error) {
	mutex.Lock()
	row := r.db.QueryRowContext(ctx, "SELECT id, location, service_1_temperature, service_2_temperature, request_count, created_at, first_request_uuid from weather_queries WHERE location = $1 LIMIT 1", location)
	mutex.Unlock()

	var weather model.DBWeather
	var flag bool
	err := row.Scan(&weather.ID, &weather.Location, &weather.Service1Temperature, &weather.Service2Temperature, &weather.RequestCount, &weather.CreatedAt, &weather.FirstRequestUUID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = r.addWeather(ctx, location)
			flag = true
			if err != nil {
				return 0, err
			}
		} else {
			return 0, err
		}
	}

	if weather.RequestCount == 0 && !flag {
		err = r.updateForFirstRequest(ctx, weather.ID)
		if err != nil {
			return 0, err
		}
	} else if !flag {
		err = r.increaseRequestCount(ctx, weather.ID, weather.RequestCount)
		if err != nil {
			return 0, err
		}
	}

	averageTemperature, err := r.listenLocationChannelForAverageTemperature(location)
	if err != nil {
		return 0, err
	}

	return averageTemperature, nil
}

func (r *repository) addWeather(ctx context.Context, location string) error {
	mutex.Lock()
	defer mutex.Unlock()
	_, err := r.db.ExecContext(ctx, "INSERT INTO weather_queries(location, request_count, created_at, first_request_uuid) VALUES ($1, 1, $2, $3)", location, time.Now().UTC(), uuid.New())
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) updateForFirstRequest(ctx context.Context, id int) error {
	mutex.Lock()
	defer mutex.Unlock()
	_, err := r.db.ExecContext(ctx, "UPDATE weather_queries SET request_count = 1, first_request_uuid = $1 WHERE id = $2", uuid.New(), id)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) increaseRequestCount(ctx context.Context, id int, requestCount uint8) error {
	mutex.Lock()
	defer mutex.Unlock()
	requestCount++
	_, err := r.db.ExecContext(ctx, "UPDATE weather_queries SET request_count = $1 WHERE id = $2", requestCount, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) listenLocationChannelForAverageTemperature(location string) (float32, error) {
	listener := pq.NewListener(r.connStr, 1*time.Second, 3*time.Second, nil)
	defer listener.Close()
	err := listener.Listen(location)
	if err != nil {
		log.Println(err)
	}

	for {
		notification := <-listener.Notify
		if notification != nil {
			f64, err := strconv.ParseFloat(notification.Extra, 64)
			if err != nil {
				return 0, err
			}

			return float32(f64), nil
		}
	}
}
