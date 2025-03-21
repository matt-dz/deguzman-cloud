CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    email_verified BOOLEAN NOT NULL DEFAULT FALSE,
    registered_2fa BOOLEAN NOT NULL DEFAULT FALSE,
    totp_key BYTEA,
    recovery_code BYTEA,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now ()
);

CREATE INDEX email_index ON users (email);

CREATE TABLE sessions (
    id TEXT NOT NULL PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    expires_at TIMESTAMPTZ NOT NULL,
    two_factor_verified BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE email_verification_request (
    id SERIAL NOT NULL PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    email TEXT NOT NULL,
    code TEXT NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE password_reset_session (
    id TEXT NOT NULL PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    email TEXT NOT NULL,
    code TEXT NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,
    email_verified BOOLEAN NOT NULL DEFAULT FALSE,
    two_factor_verified BOOLEAN NOT NULL DEFAULT FALSE
);
