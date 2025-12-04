package main

import (
	"fmt"
	"log"
	"net/http"

	"eaglebank/internal/infrastructure/implementations/config"
	databasae "eaglebank/internal/infrastructure/implementations/databasae"
)

func main() {
	cfg := config.New()
	dbase, err := databasae.New(cfg)
	if err != nil {
		fmt.Println("Bootstrap error: ", err.Error())
		return
	}

	db, err := dbase.Connect()
	if err != nil {
		fmt.Println("Database error: ", err.Error())
		return
	}

	defer db.Close()

	router, err := NewRouter(db, cfg)
	if err != nil {
		fmt.Println("Error creating routes: ", err.Error())
		return
	}

	port := cfg.GetPort()

	log.Printf("Starting server on %s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}
}
