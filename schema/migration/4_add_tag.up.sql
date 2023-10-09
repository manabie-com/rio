-- Add columns and btree indexes to the `stubs` table
ALTER TABLE rio_services.stubs
    ADD COLUMN tag VARCHAR(127) DEFAULT '';

-- Add columns and btree indexes to the `incoming_requests` table
ALTER TABLE rio_services.incoming_requests
    ADD COLUMN tag VARCHAR(127) DEFAULT '';

-- Create indexes for the `stubs` table
CREATE INDEX idx_tag_stubs ON rio_services.stubs (tag);
CREATE INDEX idx_namespace_stubs ON rio_services.stubs (namespace);

-- Create indexes for the `incoming_requests` table
CREATE INDEX idx_tag_incoming_requests ON rio_services.incoming_requests (tag);
CREATE INDEX idx_namespace_incoming_requests ON rio_services.incoming_requests (namespace);
