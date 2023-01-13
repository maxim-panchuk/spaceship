package pirate

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type PirateRepo struct {
	connection *pgx.Conn
}

func NewPirateRepo(conn *pgx.Conn) *PirateRepo {
	return &PirateRepo{
		connection: conn,
	}
}

func (r *PirateRepo) Insert(name string, power, sectorId int) (int, error) {
	var id int

	sqlString := `
	INSERT INTO space_pirate (pirate_name, pirate_power, sector_id)
	VALUES ($1, $2, $3) RETURNING pirate_id`

	row := r.connection.QueryRow(context.Background(),
		sqlString, name, power, sectorId)

	err := row.Scan(&id)

	if err != nil {
		return -1, err
	}

	return id, nil
}
