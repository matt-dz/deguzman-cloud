CREATE TABLE files (
	key uuid PRIMARY KEY default gen_random_uuid(),
    	hash text NOT NULL,
    	file_path VARCHAR(1023) NOT NULL,
    	created_at TIMESTAMPTZ NOT NULL default now(),
    	updated_at TIMESTAMPTZ NOT NULL default now(),
    	tags text[] NOT NULL default '{}'
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
	NEW.updated_at = CURRENT_TIMESTAMP;
    	RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON files
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TABLE tokens (
	token text PRIMARY KEY,
    	created_at TIMESTAMPTZ NOT NULL default now(),
    	updated_at TIMESTAMPTZ NOT NULL default now(),
    	deleted_at TIMESTAMPTZ
);

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON tokens
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TABLE heartbeats (
	id SERIAL PRIMARY KEY,
	created_at TIMESTAMPTZ NOT NULL default now()
);
