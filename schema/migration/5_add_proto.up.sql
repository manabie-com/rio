-- Create table `protos`
CREATE TABLE IF NOT EXISTS rio_services.protos (
   id BIGSERIAL PRIMARY KEY,
   name VARCHAR(255) NOT NULL DEFAULT '',
   file_id VARCHAR(63) NOT NULL DEFAULT '',
   methods JSON,
   types JSON,
   created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Add column `protocol` and indexes to the `stubs` table
ALTER TABLE rio_services.stubs
    ADD COLUMN protocol VARCHAR(31) DEFAULT 'http';
CREATE INDEX idx_protocol_stubs ON rio_services.stubs (protocol);
CREATE INDEX idx_updated_at_stubs ON rio_services.stubs (updated_at);