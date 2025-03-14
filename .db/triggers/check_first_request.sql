CREATE OR REPLACE FUNCTION check_first_request()
RETURNS TRIGGER AS $$
DECLARE 
    service_1_temp FLOAT;
    service_2_temp FLOAT;
BEGIN
  IF NEW.request_count = 1 THEN
    PERFORM pg_sleep(5);
    
    IF NEW.first_request_uuid = get_present_first_request_uuid(NEW.id) THEN
        SELECT service_1_temperature, service_2_temperature INTO service_1_temp, service_2_temp FROM get_temperatures(NEW.location);

        PERFORM notify_users(NEW.location, service_1_temp, service_2_temp);
        PERFORM update_weather(NEW.id, service_1_temp, service_2_temp);
    END IF;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

