package main

import (
	"ef_md_test/internal/config"
	"ef_md_test/internal/models"
	"fmt"

	// "ef_md_test/internal/handlers"
	"ef_md_test/internal/repositories"
	"ef_md_test/internal/services"
	"ef_md_test/pkg/parser"
	"log"
	"net/http"
)

func main() {
	psr := parser.NewParser(http.DefaultClient)

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	rep := repositories.NewRepository(db)

	ser := services.NewService(rep, psr)
	id, err := ser.Create(models.RequestDTO{Name: "Dmitriy", Surname: "Kadisheb", Pantronymic: "ASAS"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)

	p, err := ser.GetById(id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*p)
	// handler := handlers.NewHandler(ser)
}
