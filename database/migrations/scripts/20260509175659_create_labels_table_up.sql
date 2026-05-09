-- Migration: create_labels_table (UP)
-- Created at: 2026-05-09 17:56:59
-- Ref: VNC-37 — labels table + record_labels junction table
--              record_labels replaces the old labels varchar column on records

BEGIN;

CREATE TABLE labels (
    id              VARCHAR(255) DEFAULT concat('label-', uuid_generate_v4()) NOT NULL PRIMARY KEY,
    user_id         VARCHAR(255) NOT NULL REFERENCES users(id),
    name            VARCHAR(255) NOT NULL,
    is_archived     BOOLEAN NOT NULL DEFAULT FALSE,
    mark_for_delete BOOLEAN NOT NULL DEFAULT FALSE,
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at      TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE record_labels (
    record_id   VARCHAR(255) NOT NULL REFERENCES records(id),
    label_id    VARCHAR(255) NOT NULL REFERENCES labels(id),
    PRIMARY KEY (record_id, label_id)
);

COMMIT;
