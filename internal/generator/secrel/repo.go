package secrel

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

type SecrelRepository struct {
	connection *pgx.Conn
}

func NewSecrelRepository(conn *pgx.Conn) *SecrelRepository {
	return &SecrelRepository{
		connection: conn,
	}
}

func (r *SecrelRepository) Insert(sector1, sector2, distance int) (int, int, int, int, error) {
	var secRelId int

	row := r.connection.QueryRow(context.Background(),
		"INSERT INTO sec_rel (sector_id1, sector_id2, distance) VALUES ($1, $2, $3) RETURNING sec_rel_id",
		sector1, sector2, distance)

	err := row.Scan(&secRelId)

	if err != nil {
		log.Fatal(err)
		return -1, -1, -1, -1, err
	}

	return secRelId, sector1, sector2, distance, nil
}
