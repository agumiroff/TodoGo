package main

import (
	"TodoList/internal/wire"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("app.env")
	if err != nil {
		return
	}

	app, err := wire.BuildApp()
	if err != nil {
		panic(err)
	}
	app.Start()
}
