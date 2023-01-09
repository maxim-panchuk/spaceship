package repo

import (
	"context"
	"strconv"

	"github.com/jackc/pgx/v4"
)

type FactoryRepository struct {
	connection *pgx.Conn
}

func NewFactoryRepository(conn *pgx.Conn) *FactoryRepository {
	return &FactoryRepository{
		connection: conn,
	}
}

func (fr *FactoryRepository) Insert(planetId int, factoryName string) (string, error) {
	factoryId := new(int)

	row := fr.connection.QueryRow(context.Background(),
		"INSERT INTO factory (planet_id, factory_name) VALUES ($1, $2) RETURNING factory_id",
		planetId, factoryName)

	err := row.Scan(&factoryId)

	if err != nil {
		return "nil", err
	}

	return strconv.Itoa(*factoryId), nil
}
