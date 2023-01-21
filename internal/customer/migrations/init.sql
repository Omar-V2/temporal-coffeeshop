CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS customer (
        id uuid DEFAULT uuid_generate_v4 (),
        request_id uuid NOT NULL,
        first_name VARCHAR NOT NULL,
        last_name VARCHAR NOT NULL,
        email VARCHAR NOT NULL,
        phone_number VARCHAR NOT NULL,
        phone_verified BOOLEAN DEFAULT false,
        PRIMARY KEY (id)
);