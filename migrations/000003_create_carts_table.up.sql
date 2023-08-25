CREATE TABLE IF NOT EXISTS carts(
	id bigserial NOT NULL,
	entree_id bigserial NOT NULL,
	quantity int NOT NULL,
	customer_id bigserial NOT NULL,

	PRIMARY KEY(id),
	FOREIGN KEY(entree_id) REFERENCES entrees(id) ON DELETE CASCADE,
	FOREIGN KEY(customer_id) REFERENCES customers(id) ON DELETE CASCADE
);
