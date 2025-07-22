package repo

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kraxarn/website/db"
)

type UserFlags int

const (
	UserDefault UserFlags = iota
	UserAdmin
)

type Users struct {
	conn *pgxpool.Conn
}

func NewUsers(conn *pgxpool.Conn) Users {
	return Users{
		conn: conn,
	}
}

func (u Users) Insert(username string, password []byte, flags UserFlags) (db.Id, error) {
	var id db.Id

	err := u.conn.QueryRow(context.Background(), `
		insert into users (username, password, flags)
		values ($1, $2, $3)
		returning id
	`, username, password, flags).Scan(&id)

	return id, err
}
