package repo

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kraxarn/website/data"
	"github.com/kraxarn/website/db"
)

type Users struct {
	conn *pgxpool.Conn
}

func NewUsers(conn *pgxpool.Conn) Users {
	return Users{
		conn: conn,
	}
}

func (u Users) Insert(username string, password []byte, flags data.UserFlags) (db.Id, error) {
	var id db.Id

	err := u.conn.QueryRow(context.Background(), `
		insert into users (username, password, flags)
		values ($1, $2, $3)
		returning id
	`, username, password, flags).Scan(&id)

	return id, err
}

func (u Users) Password(username string) ([]byte, error) {
	var password []byte

	err := u.conn.QueryRow(context.Background(), `
		select password from users where username = $1
	`, username).Scan(&password)

	return password, err
}

func (u Users) Id(username string) (db.Id, error) {
	var id db.Id

	err := u.conn.QueryRow(context.Background(), `
		select id from users where username = $1
	`, username).Scan(&id)

	return id, err
}

func (u Users) Flags(userId db.Id) (data.UserFlags, error) {
	var flags data.UserFlags

	err := u.conn.QueryRow(context.Background(), `
		select flags from users where id = $1
	`, userId).Scan(&flags)

	return flags, err
}
