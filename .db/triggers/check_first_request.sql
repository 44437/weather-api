CREATE OR REPLACE FUNCTION check_first_request()
RETURNS TRIGGER AS $$
DECLARE
  query TEXT;
BEGIN
  query := format(
      'SELECT check_first_request_bg(%L, %L, %L, %L)',
      NEW.id,
      NEW.location,
      NEW.request_count,
      NEW.first_request_uuid
  );

  PERFORM pg_background_launch(query);
 
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

