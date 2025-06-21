package main

import (
	"log"
	"net/http"

	"ef_md_test/internal/config"
	"ef_md_test/internal/handlers"
	"ef_md_test/internal/repositories"
	"ef_md_test/internal/services"
	"ef_md_test/pkg/parser"
)

func main() {
	psr := parser.NewParser(http.DefaultClient)

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	rep := repositories.NewRepository(db)

	ser := services.NewService(rep, psr)

	handlers := handlers.NewHandler(ser)

	mux := http.NewServeMux()

	mux.Handle("/people", handlers)
	mux.Handle("/people/", handlers)

	http.ListenAndServe(":80", mux)
}
