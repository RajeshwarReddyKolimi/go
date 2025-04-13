package main

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dbName string) (*gorm.DB, error) {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	if username == "" || password == "" {
		return nil, fmt.Errorf("DB_USERNAME and DB_PASSWORD must be set")
	}
	fmt.Println(username, password)
	connectionUrl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", "localhost", username, password, dbName)
	db, err := gorm.Open(postgres.Open(connectionUrl), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to db: ", err)
		return nil, err
	}
	fmt.Println("Connected to db")
	return db, nil
}

type TestDb struct {
	db *gorm.DB
}

func NewDb(dbName string) *TestDb {
	db, err := Connect(dbName)
	if err != nil {
		return nil
	}
	return &TestDb{db}
}
