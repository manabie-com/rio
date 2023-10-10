-- Add a column `description` to the `stubs` table
ALTER TABLE public.stubs
    ADD COLUMN description VARCHAR(511) DEFAULT '';
