-- Migration: alter_labels_table_add_column_color (UP)
-- Created at: 2026-05-09 18:39:05

BEGIN;

ALTER TABLE labels ADD COLUMN color VARCHAR(255) NOT NULL;

COMMIT;
