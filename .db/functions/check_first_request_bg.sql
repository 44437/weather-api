CREATE OR REPLACE FUNCTION check_first_request_bg(location_id INT, location TEXT, request_count SMALLINT, first_request_uuid UUID)
RETURNS VOID AS $$
DECLARE 
    service_1_temp FLOAT;
    service_2_temp FLOAT;
BEGIN
  IF request_count = 1 THEN
    PERFORM pg_sleep(5);

    IF first_request_uuid = get_present_first_request_uuid(location_id) THEN
        SELECT service_1_temperature, service_2_temperature INTO service_1_temp, service_2_temp FROM get_temperatures(location);

        PERFORM notify_users(location, service_1_temp, service_2_temp);
        PERFORM update_weather(location_id, service_1_temp, service_2_temp);
    END IF;
  END IF;
END;
$$ LANGUAGE plpgsql;

