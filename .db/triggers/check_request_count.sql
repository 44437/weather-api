CREATE OR REPLACE FUNCTION check_request_count()
RETURNS TRIGGER AS $$
DECLARE
  query TEXT;
BEGIN
  query := format(
      'SELECT check_request_count_bg(%L, %L, %L)',
      NEW.id,
      NEW.location,
      NEW.request_count
  );

  PERFORM pg_background_launch(query);

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

