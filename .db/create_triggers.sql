CREATE OR REPLACE TRIGGER check_request_count_trigger
AFTER INSERT OR UPDATE ON weather_queries
FOR EACH ROW
EXECUTE FUNCTION check_request_count();

----

CREATE OR REPLACE TRIGGER check_first_request_trigger
AFTER INSERT OR UPDATE ON weather_queries
FOR EACH ROW
EXECUTE FUNCTION check_first_request();
