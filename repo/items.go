package repo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kraxarn/website/db"
)

type Items struct {
	conn *pgx.Conn
}

type EditItem struct {
	Key      string
	Username string
	Value    *string
	Icon     *string
	Priority *int
}

type Item struct {
	Key   string
	Value string
	Icon  string
}

// Could just have one NewItems for pgx.Conn, but I like this better

func NewItemsFromPool(conn *pgxpool.Conn) Items {
	return Items{
		conn: conn.Conn(),
	}
}

func NewItemsFromTx(tx pgx.Tx) Items {
	return Items{
		conn: tx.Conn(),
	}
}

func (i Items) SelectAllForEdit() ([]EditItem, error) {
	rows, err := i.conn.Query(context.Background(), `
		select texts.key, users.username, items.value, items.icon, items.priority
		from texts
			right join users on users.id = texts.editor
			left join items on texts.key = items.key
		order by items.priority desc
	`)

	if err != nil {
		return nil, err
	}

	var items []EditItem

	for rows.Next() {
		item := EditItem{}

		err = rows.Scan(&item.Key, &item.Username, &item.Value, &item.Icon, &item.Priority)
		if err != nil {
			rows.Close()
			return nil, err
		}

		items = append(items, item)
	}

	return items, err
}

func (i Items) SelectAll() ([]Item, error) {
	rows, err := i.conn.Query(context.Background(), `
		select key, value, icon
		from items
		order by priority desc
	`)

	if err != nil {
		return nil, err
	}

	var items []Item

	for rows.Next() {
		item := Item{}

		err = rows.Scan(&item.Key, &item.Value, &item.Icon)
		if err != nil {
			rows.Close()
			return nil, err
		}

		items = append(items, item)
	}

	return items, err
}

func (i Items) Insert(key, value, icon string, priority int) (db.Id, error) {
	var id db.Id

	err := i.conn.QueryRow(context.Background(), `
		insert into items (key, value, icon, priority)
		values ($1, $2, $3, $4)
		returning id
	`, key, value, icon, priority).Scan(&id)

	return id, err
}

func (i Items) DeleteAll() error {
	//goland:noinspection SqlWithoutWhere
	_, err := i.conn.Exec(context.Background(), `
		delete from items
	`)

	return err
}
