CREATE OR REPLACE FUNCTION check_request_count()
RETURNS TRIGGER AS $$
DECLARE 
    service_1_temp FLOAT;
    service_2_temp FLOAT;
BEGIN
  IF NEW.request_count = 4 THEN

    SELECT service_1_temperature, service_2_temperature INTO service_1_temp, service_2_temp FROM get_temperatures(NEW.location);

    PERFORM notify_users(NEW.location, service_1_temp, service_2_temp);
    PERFORM update_weather(NEW.id, service_1_temp, service_2_temp);
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

