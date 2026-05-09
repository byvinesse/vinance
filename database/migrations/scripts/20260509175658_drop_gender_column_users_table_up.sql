-- Migration: drop_gender_column_users_table (UP)
-- Created at: 2026-05-09 17:56:58
-- Ref: VNC-37 — gender field removed from users table

BEGIN;

ALTER TABLE users DROP COLUMN IF EXISTS gender;

COMMIT;
