-- Create the database if it doesn't exist
DO $$ 
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'telemetry') THEN
    CREATE DATABASE telemetry;
  END IF;
END $$;
-- Connect to the database
\c telemetry;

-- Create the sensors table
CREATE TABLE IF NOT EXISTS device (
    id INTEGER PRIMARY KEY,
    type VARCHAR(50) NOT NULL
);

-- -- Create indexes for common queries
-- CREATE INDEX IF NOT EXISTS idx_sensors_type ON sensors(type);
-- CREATE INDEX IF NOT EXISTS idx_sensors_location ON sensors(location);
-- CREATE INDEX IF NOT EXISTS idx_sensors_status ON sensors(status);

