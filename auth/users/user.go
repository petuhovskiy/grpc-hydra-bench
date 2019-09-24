package users

import "time"

var Schema = `
CREATE TABLE users (
	id SERIAL NOT NULL PRIMARY KEY,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	username text UNIQUE
);

INSERT INTO users (username) VALUES ('Arthur')
`

type User struct {
	ID        uint      `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Username  string    `db:"username"`
}
