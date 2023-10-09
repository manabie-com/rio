-- Add a column `proxy` to the `stubs` table
ALTER TABLE rio_services.stubs
    ADD COLUMN proxy JSON DEFAULT NULL;
