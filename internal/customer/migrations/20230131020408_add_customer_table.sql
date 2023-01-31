-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- TODO: Add request id to customer table
-- TODO: Make email and phone_number fields unique
CREATE TABLE IF NOT EXISTS customer (
        id uuid DEFAULT uuid_generate_v4 (),
        first_name VARCHAR NOT NULL,
        last_name VARCHAR NOT NULL,
        email VARCHAR NOT NULL, 
        phone_number VARCHAR NOT NULL,
        phone_verified BOOLEAN DEFAULT false,
        PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS customer;
-- +goose StatementEnd
