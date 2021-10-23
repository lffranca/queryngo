package main

import (
	"github.com/lffranca/queryngo/pkg/gmigrate"
	"github.com/lffranca/queryngo/pkg/postgres"
	"log"
	"os"
)

func main() {
	connString := os.Getenv("DB_CONN_STRING")
	db, err := postgres.New(&connString)
	if err != nil {
		log.Panicln(err)
	}

	defer db.Close()

	migrationsPath := os.Getenv("MIGRATIONS_PATH")
	client, err := gmigrate.New(db.DB(), &migrationsPath)
	if err != nil {
		log.Panicln(err)
	}

	if err := client.Postgres.Up(); err != nil {
		log.Panicln(err)
	}

	log.Println("SUCCESS")
}

func init() {
	envs := []string{
		"MIGRATIONS_PATH",
		"DB_CONN_STRING",
	}

	for _, env := range envs {
		if _, ok := os.LookupEnv(env); !ok {
			log.Panicf("env var is required: %s\n", env)
		}
	}
}
