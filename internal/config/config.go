package config

import (
	"ef_md_test/internal/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	host := os.Getenv("HOST")
	user := os.Getenv("USER_DB")
	password := os.Getenv("PASSWORD_DB")
	dbname := os.Getenv("NAME_DB")
	port := os.Getenv("PORT")

	// fmt.Printf("DB Config: host=%s user=%s password=%s dbname=%s port=%s\n",
	//     host, user, password, dbname, port)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)

	fmt.Printf("DSN: %s\n", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.Person{})
	return db, nil
}
