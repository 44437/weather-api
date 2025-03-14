CREATE OR REPLACE FUNCTION get_temperatures(location TEXT)  
RETURNS TABLE (service_1_temperature FLOAT, service_2_temperature FLOAT) AS $$  
DECLARE
    response http_response;
    temperatures JSON;
BEGIN
    response := http_get('https://weather-collector-function-281137144038.europe-north1.run.app?location=' || location);

    IF response.status = 200 THEN
        temperatures := response.content::JSON;

        service_1_temperature := (temperatures->>'service_1_temperature')::FLOAT;
        service_2_temperature := (temperatures->>'service_2_temperature')::FLOAT;
        
        RETURN QUERY SELECT service_1_temperature, service_2_temperature;
    ELSE
        RAISE EXCEPTION 'Status is not OK: %', response.status;
    END IF;

EXCEPTION WHEN others THEN
    RAISE EXCEPTION 'Something went wrong: %', SQLERRM;
END;
$$ LANGUAGE plpgsql;

