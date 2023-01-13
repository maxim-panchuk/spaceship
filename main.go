package main

import (
	"context"
	"spaceship/server"
)

func main() {
	app := server.NewApp()
	defer app.Conn.Close(context.Background())

	generate(app)

	app.Run()
}

func generate(a *server.App) {
	a.Generate()
}

// TODO:
// 1. Написать логику создания фабрик.
