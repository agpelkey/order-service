CREATE TABLE IF NOT EXISTS tokens (
	hash bytea PRIMARY KEY,
	customer_id bigint NOT NULL REFERENCES customers ON DELETE CASCADE,
	expiry timestamp(0) with time zone NOT NULL,
	scope text NOT NULL
);
