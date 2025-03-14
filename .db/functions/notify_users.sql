CREATE OR REPLACE FUNCTION notify_users(location TEXT, service_1_temperature FLOAT, service_2_temperature FLOAT)
RETURNS VOID AS $$
BEGIN
    PERFORM pg_notify(location, ((service_1_temperature + service_2_temperature) / 2)::TEXT);
END;
$$ LANGUAGE plpgsql;