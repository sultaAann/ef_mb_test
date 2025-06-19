package main

import (
	"ef_md_test/internal/service"
	"fmt"
	"net/http"
)

func main() {
	parser := service.NewParser(http.DefaultClient)
	fmt.Println(parser.Parse_age("Dmitriy"))
	fmt.Println(parser.Parse_gender("Dmitriy"))
	fmt.Println(parser.Parse_nation("Dmitriy"))
}
