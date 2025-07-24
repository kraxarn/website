package repo

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kraxarn/website/db"
)

type Texts struct {
	conn *pgxpool.Conn
}

func NewTexts(conn *pgxpool.Conn) Texts {
	return Texts{
		conn: conn,
	}
}

func (t Texts) Exists(key string) (bool, error) {
	var count int64

	err := t.conn.QueryRow(context.Background(), `
		select count(*) from texts where key = $1
	`, key).Scan(&count)

	return count > 0, err
}

func (t Texts) Insert(key, value string, userId db.Id) (db.Id, error) {
	var id db.Id

	err := t.conn.QueryRow(context.Background(), `
		insert into texts (key, value, editor, timestamp)
		values ($1, $2, $3, current_timestamp)
		returning id
	`, key, value, userId).Scan(&id)

	return id, err
}

func (t Texts) Update(key, value string, userId db.Id) (int64, error) {
	tag, err := t.conn.Exec(context.Background(), `
		update texts set value = $2 where key = $1 and editor = $3
	`, key, value, userId)

	if err != nil {
		return 0, err
	}

	return tag.RowsAffected(), nil
}

func (t Texts) Value(key string) (string, error) {
	var value string

	err := t.conn.QueryRow(context.Background(), `
		select value from texts where key = $1
	`, key).Scan(&value)

	return value, err
}
