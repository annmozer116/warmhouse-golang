-- Create the database if it doesn't exist
DO $$ 
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'device') THEN
    CREATE DATABASE device;
  END IF;
END $$;
-- Connect to the database
\c device;

-- Create the sensors table
CREATE TABLE IF NOT EXISTS device (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    type VARCHAR(50) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'inactive',
    serial_number VARCHAR(20) NOT NULL DEFAULT '000-QWERTY-000',

    location_id INTEGER NOT NULL,
    model_id INTEGER NOT NULL,
    owner_id INTEGER NOT NULL DEFAULT 0,

    last_updated TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

-- -- Create indexes for common queries
-- CREATE INDEX IF NOT EXISTS idx_sensors_type ON sensors(type);
-- CREATE INDEX IF NOT EXISTS idx_sensors_location ON sensors(location);
-- CREATE INDEX IF NOT EXISTS idx_sensors_status ON sensors(status);

-- Добавим тестовые данные
INSERT INTO device (name, type, location_id, owner_id, serial_number, model_id) VALUES
('Гостиный термостат', 'temperature', 1, 1, '000-QWERTY-000', 1),
('Детектор дыма на кухне', 'smoke_detector', 2, 1, '000-QWERTY-001', 4),
('Уличная камера', 'camera', 3, 1, '000-QWERTY-002', 6),
('Спальня свет', 'lighting', 4, 1, '000-QWERTY-003', 9),
('Люстра на кухне', 'lighting', 2, 1, '000-QWERTY-004',9);
