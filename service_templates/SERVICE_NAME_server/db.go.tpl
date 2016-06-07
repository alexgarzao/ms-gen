package main

import (
	"fmt"
	"log"

	_ "github.com/alexgarzao/ms-gen/code/{{.ServiceName}}_common"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type DB struct {
	gorm.DB
}

func (i *DB) InitDB() {

	// Check config to connect into DB
	dbHost := viper.GetString("db_host")
	if dbHost == "" {
		log.Fatalf("Config db_host not found!")
	}

	dbName := viper.GetString("db_name")
	if dbName == "" {
		log.Fatalf("Config db_name not found!")
	}

	dbUser := viper.GetString("db_user")
	if dbUser == "" {
		log.Fatalf("Config db_user not found!")
	}

	dbPassword := viper.GetString("db_password")
	if dbPassword == "" {
		log.Fatalf("Config db_password not found!")
	}

	connectionString := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable", dbHost, dbName, dbUser, dbPassword)

	var err error
	i.DB, err = gorm.Open("postgres", connectionString)

	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}

	i.DB.LogMode(true)
}

func (i *DB) InitSchema() {
	i.DB.AutoMigrate(&Model1{})
	i.DB.AutoMigrate(&Model2{})
}
