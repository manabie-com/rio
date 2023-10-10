-- Add columns and btree indexes to the `stubs` table
ALTER TABLE public.stubs
    ADD COLUMN tag VARCHAR(127) DEFAULT '';

-- Add columns and btree indexes to the `incoming_requests` table
ALTER TABLE public.incoming_requests
    ADD COLUMN tag VARCHAR(127) DEFAULT '';

-- Create indexes for the `stubs` table
CREATE INDEX idx_tag_stubs ON public.stubs (tag);
CREATE INDEX idx_namespace_stubs ON public.stubs (namespace);

-- Create indexes for the `incoming_requests` table
CREATE INDEX idx_tag_incoming_requests ON public.incoming_requests (tag);
CREATE INDEX idx_namespace_incoming_requests ON public.incoming_requests (namespace);
