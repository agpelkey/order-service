CREATE TABLE IF NOT EXISTS customers(
	id bigserial NOT NULL,
	username text NOT NULL UNIQUE,
	email text NOT NULL UNIQUE,
	password bytea NOT NULL,
	created_at timestamptz DEFAULT NOW(),
	updated_at timestamptz DEFAULT NOW(),

	PRIMARY KEY(id)
);
