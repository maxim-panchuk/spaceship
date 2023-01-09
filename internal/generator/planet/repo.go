package planet

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

type PlanetRepository struct {
	connection *pgx.Conn
}

func NewPlanetRepository(conn *pgx.Conn) *PlanetRepository {
	return &PlanetRepository{
		connection: conn,
	}
}

// вернет planet_id, planet_name, sector_id, error
func (r *PlanetRepository) Insert(planetName string, sectorId int) (int, string, int, error) {
	var planetId int

	row := r.connection.QueryRow(context.Background(),
		"INSERT INTO planet (planet_name, sector_id) VALUES ($1, $2) RETURNING planet_id",
		planetName, sectorId)

	err := row.Scan(&planetId)

	if err != nil {
		return -1, "", -1, err
	}

	return planetId, planetName, sectorId, nil
}

func (r *PlanetRepository) InsertRelation(sectorId1, sectorId2, distance int) (int, int, int, int, error) {
	var relId int
	row := r.connection.QueryRow(context.Background(),
		"INSERT INTO sec_rel (sector_id1, sector_id2, distance) VALUES ($1, $2, $3) RETURNING sec_rel_id",
		sectorId1, sectorId2, distance)

	err := row.Scan(&relId)

	if err != nil {
		log.Fatal(err)
		return -1, -1, -1, -1, err
	}

	return relId, sectorId1, sectorId2, distance, nil
}
