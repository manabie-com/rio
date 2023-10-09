-- PostgreSQL Script
-- Sat Aug  7 13:27:47 2021
-- Model: New Model    Version: 1.0

-- Disable foreign key and unique checks temporarily
SET session_replication_role = replica;

-- Drop the schema if it exists (Note: PostgreSQL does not require this step)
-- DROP SCHEMA IF EXISTS rio_services;

-- Create the schema if it does not exist
CREATE SCHEMA IF NOT EXISTS rio_services;

-- Switch to the rio_services schema (Note: PostgreSQL does not require this step)
-- SET search_path = rio_services;

-- Table `stubs`
CREATE TABLE IF NOT EXISTS rio_services.stubs (
  id BIGSERIAL PRIMARY KEY,
  namespace VARCHAR(255) NOT NULL DEFAULT '',
  request JSON,
  response JSON,
  weight INT NOT NULL DEFAULT 0,
  active INT NOT NULL DEFAULT 0,
  settings JSON,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Table `incoming_requests`
CREATE TABLE IF NOT EXISTS rio_services.incoming_requests (
  id BIGSERIAL PRIMARY KEY,
  namespace VARCHAR(255) NOT NULL DEFAULT '',
  url TEXT NOT NULL,
  method VARCHAR(31) NOT NULL DEFAULT '',
  header JSON,
  body BYTEA,
  stub_id BIGINT NOT NULL DEFAULT 0,
  curl TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Re-enable foreign key and unique checks
SET session_replication_role = DEFAULT;
