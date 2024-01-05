-- +goose Up
CREATE TABLE items (
	id UUID PRIMARY KEY,
	created_at TIMESTAMP NOT NULL CHECK (created_at <= NOW()),
	updated_at TIMESTAMP NOT NULL CHECK (updated_at >= created_at),
	name VARCHAR(72) NOT NULL,
	description VARCHAR(720) NOT NULL,
	price INT NOT NULL CHECK (price >= 0),
	sold BOOLEAN NOT NULL,
	seller_id UUID NOT NULL,
	FOREIGN KEY (seller_id) REFERENCES users(id)
);

-- +goose Down
DROP TABLE items;
