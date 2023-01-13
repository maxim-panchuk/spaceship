package carrier

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type CarrierRepository struct {
	connection *pgx.Conn
}

func NewCarrierRepository(conn *pgx.Conn) *CarrierRepository {
	return &CarrierRepository{
		connection: conn,
	}
}

func (r *CarrierRepository) Insert(name string, power int, speed int) (int, error) {
	var id int

	sqlString := `
	INSERT INTO carrier (carrier_name, carrier_power, carrier_speed)
	VALUES ($1, $2, $3) RETURNING carrier_id`

	row := r.connection.QueryRow(context.Background(),
		sqlString, name, power, speed)

	err := row.Scan(&id)

	if err != nil {
		return -1, err
	}

	return id, nil
}
