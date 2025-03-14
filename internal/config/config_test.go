package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	t.Run("When the config file cannot be reachable", func(t *testing.T) {
		path := "../../testdata"
		name := "wrong_config"

		_, err := New(path, name)
		assert.NotNil(t, err)
	})

	t.Run("When the config is created successfully", func(t *testing.T) {
		path := "../../testdata"
		name := "test_config"

		actualConfig, _ := New(path, name)

		expectedConfig := &Config{
			Server: Server{
				Port: 8080,
			},
			Postgres: Postgres{
				Host:     "postgresql",
				Port:     5432,
				User:     "admin",
				Password: "admin",
				Name:     "weather",
			},
		}

		assert.Equal(t, expectedConfig, actualConfig)
	})
}

func TestPrint(t *testing.T) {
	c := Config{}
	c.Print()
}
