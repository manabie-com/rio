-- PostgreSQL Script
-- Created on Tue Jun 21 09:28:09 2022
-- Model: New Model    Version: 1.0

-- Disable foreign key and unique checks temporarily
SET session_replication_role = replica;

-- Create the schema if not exists
CREATE SCHEMA IF NOT EXISTS rio_services;

-- Switch to the rio_services schema
SET search_path = rio_services;

-- Table `stubs`
CREATE TABLE IF NOT EXISTS stubs (
  id SERIAL PRIMARY KEY,
  namespace VARCHAR(255) NOT NULL DEFAULT '',
  request JSON,
  response JSON,
  weight INT NOT NULL DEFAULT 0,
  active INT NOT NULL DEFAULT 0,
  proxy JSON,
  settings JSON,
  description VARCHAR(511) NOT NULL DEFAULT '',
  tag VARCHAR(127) DEFAULT '',
  protocol VARCHAR(31) DEFAULT 'http',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
  );

-- Indexes for `stubs` table
CREATE INDEX IF NOT EXISTS idx_tag ON stubs (tag);
CREATE INDEX IF NOT EXISTS idx_namespace ON stubs (namespace);
CREATE INDEX IF NOT EXISTS idx_protocol ON stubs (protocol);
CREATE INDEX IF NOT EXISTS idx_updated_at ON stubs (updated_at);

-- Table `incoming_requests`
CREATE TABLE IF NOT EXISTS incoming_requests (
  id SERIAL PRIMARY KEY,
  namespace VARCHAR(255) NOT NULL DEFAULT '',
  tag VARCHAR(127) DEFAULT '',
  url TEXT NOT NULL,
  "method" VARCHAR(31) NOT NULL DEFAULT '',
  header JSON,
  body BYTEA,
  stub_id INT NOT NULL DEFAULT 0,
  curl TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
  );

-- Indexes for `incoming_requests` table
CREATE INDEX IF NOT EXISTS idx_tag ON incoming_requests (tag);
CREATE INDEX IF NOT EXISTS idx_namespace ON incoming_requests (namespace);

-- Table `protos`
CREATE TABLE IF NOT EXISTS protos (
  id SERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL DEFAULT '',
  file_id VARCHAR(63) NOT NULL DEFAULT '',
  methods JSON,
  types JSON,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
  );

-- Index for `protos` table
CREATE INDEX IF NOT EXISTS idx_updated_at ON protos (updated_at);

-- Re-enable foreign key and unique checks
SET session_replication_role = DEFAULT;
