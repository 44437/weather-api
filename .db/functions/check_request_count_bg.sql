CREATE OR REPLACE FUNCTION check_request_count_bg(location_ID INT, location TEXT, request_count SMALLINT)
RETURNS VOID AS $$
DECLARE 
    service_1_temp FLOAT;
    service_2_temp FLOAT;
BEGIN
  IF request_count = 10 THEN
    SELECT service_1_temperature, service_2_temperature INTO service_1_temp, service_2_temp FROM get_temperatures(location);

    PERFORM notify_users(location, service_1_temp, service_2_temp);
    PERFORM update_weather(location_ID, service_1_temp, service_2_temp);
  END IF;
END;
$$ LANGUAGE plpgsql;

