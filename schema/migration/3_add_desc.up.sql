-- Add a column `description` to the `stubs` table
ALTER TABLE rio_services.stubs
    ADD COLUMN description VARCHAR(511) DEFAULT '';
