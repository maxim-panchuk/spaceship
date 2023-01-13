package server

import (
	"context"
	"fmt"
	"os"
	"spaceship/auth"
	"spaceship/auth/delivery"
	"spaceship/auth/repo"
	"spaceship/auth/usecase"
	"spaceship/internal/deal"
	"spaceship/internal/factory"
	d "spaceship/internal/factory/delivery"
	fr "spaceship/internal/factory/repo"
	fu "spaceship/internal/factory/usecase"
	"spaceship/internal/item"
	di "spaceship/internal/item/delivery"

	ir "spaceship/internal/item/repo"
	iu "spaceship/internal/item/usecase"

	dd "spaceship/internal/deal/delivery"
	dr "spaceship/internal/deal/repo"
	du "spaceship/internal/deal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type App struct {
	authUseCase    auth.UseCase
	factoryUseCase factory.UseCase
	itemUseCase    item.UseCase
	dealUseCase    deal.UseCase
	Conn           *pgx.Conn
}

func NewApp() *App {
	conn := initDB()

	userRepo := repo.NewUserRepository(conn)
	factoryRepo := fr.NewFactoryRepository(conn)
	itemRepo := ir.NewItemRepository(conn)
	dealRepo := dr.NewFactoryRepository(conn)

	authUseCase := usecase.NewAuthorizer(userRepo, "hash_salt")
	factoryUseCase := fu.NewFactoryCreator(factoryRepo)
	itemUseCase := iu.NewItemUseCase(itemRepo)
	dealUseCase := du.NewDealUseCase(dealRepo)

	return &App{
		authUseCase:    authUseCase,
		factoryUseCase: factoryUseCase,
		itemUseCase:    itemUseCase,
		dealUseCase:    dealUseCase,
		Conn:           conn,
	}

}

func (a *App) Run() {
	router := gin.Default()

	api := router.Group("/")
	app := router.Group("/app")

	app.Use(delivery.TokenAuthMiddleware())

	d.RegisterHTTPEndpoints(app, a.factoryUseCase)
	di.RegisterHTTPEndpoints(app, a.itemUseCase)
	dd.RegisterHTTPEndpoints(app, a.dealUseCase)

	delivery.RegisterHTTPEndpoints(api, a.authUseCase)

	router.Run("localhost:8080")
}

func initDB() *pgx.Conn {
	dbUrl := "postgres://maksim:kilibok47@localhost:5432/spaceship"
	conn, err := pgx.Connect(context.Background(), dbUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unnable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}

func (a *App) Generate() {
	// sectorRepo := sector.NewSectorRepository(a.Conn)
	// sectorGenerator := sector.NewSectorGenerator(*sectorRepo)

	// planetRepo := planet.NewPlanetRepository(a.Conn)
	// planetGen := planet.NewPlanetGenerator(*planetRepo, *sectorRepo)

	// relRepo := secrel.NewSecrelRepository(a.Conn)
	// relGen := secrel.NewSecrelGenerator(*relRepo, *sectorRepo)

	// // sectorGenerator.Generate()
	// // planetGen.Generate()
	// // relGen.Generate()

	// pirateRepo := pirate.NewPirateRepo(a.Conn)
	// pirateGen := pirate.NewPirateGen(*pirateRepo, *sectorRepo)

	// pirateGen.Generate()

	// carrierRepo := carrier.NewCarrierRepository(a.Conn)
	// cg := carrier.NewCarrierGen(*carrierRepo)
	// cg.Generate()
}
