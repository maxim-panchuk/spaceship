package repo

import (
	"context"
	"fmt"
	"log"
	"os"
	"spaceship/entity"

	"github.com/jackc/pgx/v4"
)

type DealRepository struct {
	connection *pgx.Conn
}

func NewFactoryRepository(conn *pgx.Conn) *DealRepository {
	return &DealRepository{
		connection: conn,
	}
}

// Добавить требование на поставку
func (r *DealRepository) InsertRequire(factoryBuyerId, factorySellerId, itemId, amount int) (int, error) {
	sqlStatement := `
	INSERT INTO delivery_require_deal (factory_buyer_id, factory_seller_id, item_id, amount)
	VALUES ($1, $2, $3, $4) RETURNING dlvr_req_prdct_rel_id`

	var id int

	err := r.connection.QueryRow(context.Background(), sqlStatement, factoryBuyerId, factorySellerId, itemId, amount).Scan(&id)

	if err != nil {
		return -1, err
	}

	return id, err
}

func (r *DealRepository) GetRequiresBySeller(factorySellerId int) ([]entity.DeliveryRequireDeal, error) {
	dealReqSlice := make([]entity.DeliveryRequireDeal, 0)

	sqlStatement := `
	SELECT * 
	FROM delivery_require_deal
	WHERE factory_seller = $1`

	rows, err := r.connection.Query(context.Background(), sqlStatement, factorySellerId)

	if err != nil {
		log.Fatalf(fmt.Sprintf("Error while getting factories where item is: %v", err))
		return nil, err
	}

	for rows.Next() {
		var f entity.DeliveryRequireDeal
		err := rows.Scan(&f.DlvrReqPrdctRelId, &f.FactoryBuyerId, &f.FactorySellerId, &f.ItemId, &f.Amount)
		if err != nil {
			log.Fatalf("Error while scanning requires: %v", err)
			return nil, err
		}
		dealReqSlice = append(dealReqSlice, f)
	}

	if rows.Err() != nil {
		fmt.Fprintf(os.Stderr, "rows Error: %v\n", rows.Err())
		return nil, err
	}

	defer rows.Close()

	return dealReqSlice, nil
}

func (r *DealRepository) InsertAgreement(carrierId, deliveryPrice, dealPrice int) error {
	sqlStatement := `
	INSERT INTO delivery_agreement (carrier_id, delivery_price, deal_price)
	VALUES ($1, $2, $3)`

	rows, err := r.connection.Query(context.Background(), sqlStatement, carrierId, dealPrice, dealPrice)

	if err != nil {
		log.Fatalf(fmt.Sprintf("Error while inserting agreement: %v", err))
		return err
	}

	if rows.Err() != nil {
		fmt.Fprintf(os.Stderr, "rows Error: %v\n", rows.Err())
		return err
	}

	defer rows.Close()

	return nil
}

func (r *DealRepository) GetSectorByFId(factoryId int) (entity.Sector, error) {
	sqlStatement := `
	SELECT s.sector_id, s.pirate_scale, s.sector_name
	FROM factory f
	JOIN planet p
	ON f.planet_id = p.planet_id
	JOIN sector s
	ON s.sector_id = p.sector_id
	WHERE f.factory_id = $1`

	rows, err := r.connection.Query(context.Background(), sqlStatement, factoryId)

	if err != nil {
		return entity.Sector{}, err
	}

	var f entity.Sector

	for rows.Next() {
		err := rows.Scan(&f.SectorId, &f.PirateScale, &f.SectorName)
		if err != nil {
			return entity.Sector{}, err
		}
	}

	if rows.Err() != nil {
		fmt.Fprintf(os.Stderr, "rows Error: %v\n", rows.Err())
		return entity.Sector{}, err
	}

	defer rows.Close()

	return f, nil
}

func (r *DealRepository) GetRequireById(dlvReqId int) (entity.DeliveryRequireDeal, error) {
	sqlStatement := `
	SELECT *
	FROM delivery_require_deal
	WHERE dlvr_req_prdct_rel_id = $1`

	var f entity.DeliveryRequireDeal

	row := r.connection.QueryRow(context.Background(),
		sqlStatement, dlvReqId)

	err := row.Scan(&f.DlvrReqPrdctRelId, &f.FactoryBuyerId, &f.FactorySellerId, &f.ItemId, &f.Amount)

	if err != nil {
		return entity.DeliveryRequireDeal{}, err
	}

	return f, nil
}

func (r *DealRepository) GetAllSecRel() ([]entity.SectorRelation, error) {

	sqlStatement := `
	SELECT *
	FROM sec_rel`

	secRelSlice := make([]entity.SectorRelation, 0)

	rows, err := r.connection.Query(context.Background(), sqlStatement)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		var f entity.SectorRelation

		err := rows.Scan(&f.SecRelId, &f.SectorId1, &f.SectorId2, &f.Distance)

		if err != nil {
			return nil, err
		}

		secRelSlice = append(secRelSlice, f)
	}

	if rows.Err() != nil {
		fmt.Fprintf(os.Stderr, "rows Error: %v\n", rows.Err())
		return nil, err
	}

	defer rows.Close()

	return secRelSlice, nil
}
