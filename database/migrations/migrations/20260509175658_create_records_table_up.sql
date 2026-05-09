-- Migration: create_records_table (UP)
-- Created at: 2026-05-09 17:56:58
-- Ref: VNC-37 — records table with dual-currency support (amount + base_amount)
--              labels are stored separately in record_labels (see create_labels_table migration)

BEGIN;

CREATE TABLE records (
    id              VARCHAR(255) DEFAULT concat('record-', uuid_generate_v4()) NOT NULL PRIMARY KEY,
    user_id         VARCHAR(255) NOT NULL REFERENCES users(id),
    account_id      VARCHAR(255) NOT NULL REFERENCES accounts(id),
    subcategory_id  VARCHAR(255) NOT NULL REFERENCES subcategories(id),
    amount          NUMERIC(18,6) NOT NULL,
    currency        VARCHAR(255) NOT NULL,
    base_amount     NUMERIC(18,6) NOT NULL,
    type            VARCHAR(255) NOT NULL,
    name            VARCHAR(255),
    payee           VARCHAR(255),
    payment_type    VARCHAR(255),
    payment_status  VARCHAR(255) NOT NULL,
    is_excluded     BOOLEAN NOT NULL DEFAULT FALSE,
    recorded_at     TIMESTAMP WITH TIME ZONE,
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at      TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

COMMIT;
