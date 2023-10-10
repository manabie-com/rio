-- Add a column `proxy` to the `stubs` table
ALTER TABLE public.stubs
    ADD COLUMN proxy JSON DEFAULT NULL;
