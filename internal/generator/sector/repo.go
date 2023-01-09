package sector

import (
	"context"
	"log"
	"spaceship/entity"

	"github.com/jackc/pgx/v4"
)

type SectorRepository struct {
	connection *pgx.Conn
}

func NewSectorRepository(conn *pgx.Conn) *SectorRepository {
	return &SectorRepository{
		connection: conn,
	}
}

// Insert вернет sector_id, pirate_scale, sector_name, error
func (r *SectorRepository) Insert(pirateScale float32, sectorName string) (int, float32, string, error) {
	sectorId := new(int)

	row := r.connection.QueryRow(context.Background(),
		"INSERT INTO sector (pirate_scale, sector_name) VALUES ($1, $2) RETURNING sector_id",
		pirateScale, sectorName)

	err := row.Scan(&sectorId)

	if err != nil {
		return -1, -0.1, "", err
	}

	return *sectorId, pirateScale, sectorName, nil
}

// Функция вернет все секторы которые были созданы генератором секторов.
func (r *SectorRepository) GetSectors() ([]entity.Sector, error) {
	rows, err := r.connection.Query(context.Background(), "SELECT * FROM sector")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer rows.Close()

	var sectorSlice []entity.Sector

	for rows.Next() {
		sector := new(entity.Sector)

		err := rows.Scan(&sector.SectorId, &sector.PirateScale, &sector.SectorName)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		sectorSlice = append(sectorSlice, *sector)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return sectorSlice, nil
}
