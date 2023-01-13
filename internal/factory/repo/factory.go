package repo

import (
	"context"
	"fmt"
	"log"
	"os"
	"spaceship/entity"
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

func (fr *FactoryRepository) Insert(planetId int, factoryName, username string) (string, error) {
	factoryId := new(int)

	row := fr.connection.QueryRow(context.Background(),
		"INSERT INTO factory (planet_id, factory_name, username) VALUES ($1, $2, $3) RETURNING factory_id",
		planetId, factoryName, username)

	err := row.Scan(&factoryId)

	if err != nil {
		return "nil", err
	}

	// fr.connection.Close(context.Background())

	return strconv.Itoa(*factoryId), nil

}

func (fr *FactoryRepository) GetAll(username string) ([]entity.Factory, error) {

	factorySlice := make([]entity.Factory, 0)

	rows, err := fr.connection.Query(context.Background(),
		"SELECT * FROM factory WHERE username=$1", username)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var f entity.Factory
		err := rows.Scan(&f.FactoryId, &f.PlanedId, &f.FactoryName, &f.Username)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to scan %v\n", err)
			os.Exit(1)
		}

		factorySlice = append(factorySlice, f)
	}

	if rows.Err() != nil {
		fmt.Fprintf(os.Stderr, "rows Error: %v\n", rows.Err())
		os.Exit(1)
	}

	defer rows.Close()

	return factorySlice, nil
}

func (fr *FactoryRepository) InsertItemStock(factoryId, itemId int, itemPrice float32, itemAmoountStock int) error {
	rows, err := fr.connection.Query(context.Background(),
		"INSERT INTO item_factory_production (factory_id, item_id, item_price_to_sell, item_amount_instock) VALUES ($1, $2, $3, $4)",
		factoryId, itemId, itemPrice, itemAmoountStock)

	if rows.Err() != nil {
		fmt.Fprintf(os.Stderr, "rows Error: %v\n", rows.Err())
		return err
	}

	defer rows.Close()

	if err != nil {
		return err
	}

	// fr.connection.Close(context.Background())

	return nil
}

func (fr *FactoryRepository) GetFactoryStock(factoryId int) ([]entity.ItemFactoryProduction, error) {
	factoryStockSlice := make([]entity.ItemFactoryProduction, 0)

	rows, err := fr.connection.Query(context.Background(),
		"SELECT * FROM item_factory_productionn WHERE factory_id=$1", factoryId)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var f entity.ItemFactoryProduction
		err := rows.Scan(&f.ItemFactoryRelId, &f.FactoryId, &f.ItemId, &f.ItemPriceToSell, &f.ItemAmountInStock)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to scan %v\n", err)
			return nil, err
		}

		factoryStockSlice = append(factoryStockSlice, f)
	}

	if rows.Err() != nil {
		fmt.Fprintf(os.Stderr, "rows Error: %v\n", rows.Err())
		return nil, err
	}

	defer rows.Close()

	// fr.connection.Close(context.Background())

	return factoryStockSlice, nil
}

func (fr *FactoryRepository) UpdateFactoryStockAmount(factoryId, itemId, itemAmount int) error {
	sqlStatement := `
	UPDATE item_factory_production
	SET item_amount_instock = $1
	WHERE factory_id = $2
	AND item_id = $3`

	_, err := fr.connection.Exec(context.Background(), sqlStatement, itemAmount, factoryId, itemId)

	if err != nil {
		return err
	}

	// fr.connection.Close(context.Background())

	return nil
}

func (fr *FactoryRepository) GetFactoriesWhereItem(itemId int) ([]entity.Factory, error) {

	factorySlice := make([]entity.Factory, 0)

	sqlStatement := `
	SELECT f.factory_id, f.planet_id, f.username, f.factory_name
	FROM item_factory_production if
	JOIN factory f
	ON f.factory_id = if.factory_id
	WHERE if.item_id = $1`

	rows, err := fr.connection.Query(context.Background(), sqlStatement, itemId)

	if err != nil {
		log.Fatalf(fmt.Sprintf("Error while getting factories where item is: %v", err))
		return nil, err
	}

	for rows.Next() {
		var f entity.Factory
		err := rows.Scan(&f.FactoryId, &f.PlanedId, &f.Username, &f.FactoryName)
		if err != nil {
			log.Fatalf("Error while scanning factories where item is: %v", err)
			return nil, err
		}
		factorySlice = append(factorySlice, f)
	}

	if rows.Err() != nil {
		fmt.Fprintf(os.Stderr, "rows Error: %v\n", rows.Err())
		return nil, err
	}

	defer rows.Close()

	// fr.connection.Close(context.Background())

	return factorySlice, nil
}
