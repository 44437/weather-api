CREATE OR REPLACE FUNCTION get_present_first_request_uuid(location_id INT)
RETURNS UUID AS $$
DECLARE
    uuid UUID;
BEGIN
    SELECT first_request_uuid INTO uuid
    FROM weather_queries
    WHERE id = location_id;

    RETURN uuid;
END;
$$ LANGUAGE plpgsql;