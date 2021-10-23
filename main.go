package main

import (
	"github.com/lffranca/queryngo/pkg/gaws"
	"github.com/lffranca/queryngo/pkg/ginserver"
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

	awsBucket := os.Getenv("AWS_BUCKET")
	awsClient, err := gaws.New(&awsBucket)
	if err != nil {
		log.Panicln(err)
	}

	port := os.Getenv("API_PORT")
	server, err := ginserver.New(db, awsClient, &port)
	if err != nil {
		log.Panicln(err)
	}

	if err := server.Run(); err != nil {
		log.Panicln(err)
	}
}

func init() {
	envs := []string{
		"AWS_BUCKET",
		"API_PORT",
		"DB_CONN_STRING",
	}

	for _, env := range envs {
		if _, ok := os.LookupEnv(env); !ok {
			log.Panicf("env var is required: %s\n", env)
		}
	}
}
