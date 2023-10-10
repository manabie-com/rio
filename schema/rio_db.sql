-- PostgreSQL Script
-- Created on Tue Jun 21 09:28:09 2022
-- Model: New Model    Version: 1.0

-- Disable foreign key and unique checks temporarily
SET session_replication_role = replica;

-- Create the schema if not exists
CREATE SCHEMA IF NOT EXISTS public;

-- Table `stubs`
CREATE TABLE IF NOT EXISTS public.stubs (
  id SERIAL PRIMARY KEY,
  namespace VARCHAR(255) NOT NULL DEFAULT '',
  request JSON,
  response JSON,
  weight INT NOT NULL DEFAULT 0,
  active BOOL DEFAULT FALSE,
  proxy JSON,
  settings JSON,
  description VARCHAR(511) NOT NULL DEFAULT '',
  tag VARCHAR(127) DEFAULT '',
  protocol VARCHAR(31) DEFAULT 'http',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
  );

-- Indexes for `stubs` table
CREATE INDEX IF NOT EXISTS idx_tag ON public.stubs (tag);
CREATE INDEX IF NOT EXISTS idx_namespace ON public.stubs (namespace);
CREATE INDEX IF NOT EXISTS idx_protocol ON public.stubs (protocol);
CREATE INDEX IF NOT EXISTS idx_updated_at ON public.stubs (updated_at);

-- Table `incoming_requests`
CREATE TABLE IF NOT EXISTS public.incoming_requests (
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
CREATE INDEX IF NOT EXISTS idx_tag ON public.incoming_requests (tag);
CREATE INDEX IF NOT EXISTS idx_namespace ON public.incoming_requests (namespace);

-- Table `protos`
CREATE TABLE IF NOT EXISTS public.protos (
  id SERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL DEFAULT '',
  file_id VARCHAR(63) NOT NULL DEFAULT '',
  methods JSON,
  types JSON,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
  );

-- Index for `protos` table
CREATE INDEX IF NOT EXISTS idx_updated_at ON public.protos (updated_at);

-- Re-enable foreign key and unique checks
SET session_replication_role = DEFAULT;
