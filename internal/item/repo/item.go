package repo

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type ItemRepository struct {
	connection *pgx.Conn
}

func NewItemRepository(conn *pgx.Conn) *ItemRepository {
	return &ItemRepository{
		connection: conn,
	}
}

func (fr *ItemRepository) Insert(factoryId int, itemName string) (int, error) {
	itemId := new(int)

	row := fr.connection.QueryRow(context.Background(),
		"INSERT INTO item (factory_id, item_name) VALUES ($1, $2) RETURNING item_id",
		factoryId, itemName)

	err := row.Scan(&itemId)

	if err != nil {
		return -1, err
	}

	return *itemId, nil
}
