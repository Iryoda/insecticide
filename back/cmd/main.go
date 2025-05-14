package main

import (
	"log"
	"os"

	"github.com/iryoda/insecticide/src/api"
	"github.com/iryoda/insecticide/src/utils"
	"github.com/jmoiron/sqlx"
)

func main() {
	utils.InitDotEnv()
	dbUrl := os.Getenv("DATABASE_URL")

	db, err := sqlx.Connect("pgx", dbUrl)
	if err != nil {
		log.Fatal("Error connecting to the database")
	}
	defer db.Close()

	e := api.NewServer(db)

	e.Logger.Fatal(e.Start(":8080"))
}
