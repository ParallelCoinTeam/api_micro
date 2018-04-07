package main

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateConnection() (*gorm.DB, error) {
	//return gorm.Open("postgres", "postgres://khalid:khalid@127.0.0.1:5432/meem")
	//return gorm.Open("postgres", "postgres://"+os.Getenv("SECRET_USERNAME")+":"+os.Getenv("SECRET_PASSWORD")+"@database-service1:5432/meem")
	return gorm.Open("postgres", "postgres://"+os.Getenv("SECRET_USERNAME")+":"+os.Getenv("SECRET_PASSWORD")+"@meem-db-instance.c8qfd5ux55rw.us-east-2.rds.amazonaws.com:5432/meem")
}
