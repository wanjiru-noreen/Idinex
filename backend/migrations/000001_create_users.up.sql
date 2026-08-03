CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    username VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password_hash TEXT NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT users_username_unique UNIQUE (username),
    CONSTRAINT users_email_unique UNIQUE (email),

    CONSTRAINT users_username_length
        CHECK (char_length(trim(username)) >= 3),

    CONSTRAINT users_email_not_empty
        CHECK (char_length(trim(email)) > 0),

    CONSTRAINT users_password_hash_not_empty
        CHECK (char_length(password_hash) > 0)
);

COMMENT ON TABLE users IS 'Application users';