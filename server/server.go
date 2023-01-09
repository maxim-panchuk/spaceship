package server

import (
	"context"
	"fmt"
	"os"
	"spaceship/auth"
	"spaceship/auth/delivery"
	"spaceship/auth/repo"
	"spaceship/auth/usecase"
	"spaceship/internal/factory"
	d "spaceship/internal/factory/delivery"
	fr "spaceship/internal/factory/repo"
	fu "spaceship/internal/factory/usecase"

	"spaceship/internal/generator/planet"
	"spaceship/internal/generator/secrel"
	"spaceship/internal/generator/sector"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type App struct {
	authUseCase    auth.UseCase
	factoryUseCase factory.UseCase
	Conn           *pgx.Conn
}

func NewApp() *App {
	conn := initDB()

	userRepo := repo.NewUserRepository(conn)
	factoryRepo := fr.NewFactoryRepository(conn)

	authUseCase := usecase.NewAuthorizer(userRepo, "hash_salt")
	factoryUseCase := fu.NewFactoryCreator(factoryRepo)

	return &App{
		authUseCase:    authUseCase,
		factoryUseCase: factoryUseCase,
		Conn:           conn,
	}

}

func (a *App) Run() {
	router := gin.Default()

	api := router.Group("/")
	app := router.Group("/app")

	app.Use(delivery.TokenAuthMiddleware())

	d.RegisterHTTPEndpoints(app, a.factoryUseCase)
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
	sectorRepo := sector.NewSectorRepository(a.Conn)
	sectorGenerator := sector.NewSectorGenerator(*sectorRepo)

	planetRepo := planet.NewPlanetRepository(a.Conn)
	planetGen := planet.NewPlanetGenerator(*planetRepo, *sectorRepo)

	relRepo := secrel.NewSecrelRepository(a.Conn)
	relGen := secrel.NewSecrelGenerator(*relRepo, *sectorRepo)

	sectorGenerator.Generate()
	planetGen.Generate()
	relGen.Generate()

}
