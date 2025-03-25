CREATE OR REPLACE FUNCTION update_weather(location_id INTEGER, service_1_temp FLOAT, service_2_temp FLOAT)
RETURNS VOID AS $$
BEGIN
    UPDATE weather_queries
    SET 
        service_1_temperature = service_1_temp,
        service_2_temperature = service_2_temp,
        request_count = 0,
        first_request_uuid = '00000000-0000-0000-0000-000000000000'
    WHERE id = location_id;
END;
$$ LANGUAGE plpgsql;