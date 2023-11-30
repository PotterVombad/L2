package main

import (
	"dev11/internal/api"
	"dev11/internal/db/cashe"

	log "github.com/sirupsen/logrus"
)

func main() {
	host := "127.0.0.1"
	port := "8080"
	db := cashe.NewCashe()
	api := api.New(db)
	err := api.Run(host, port)
	if err != nil {
		log.Println(err)
	}
}