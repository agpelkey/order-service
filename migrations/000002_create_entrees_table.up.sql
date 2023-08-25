CREATE TABLE IF NOT EXISTS entrees(
	id bigserial NOT NULL,
	name text NOT NULL,
	description text NOT NULL,
	cost int NOT NULL,
	quantity int NOT NULL,

	PRIMARY KEY(id)

);
