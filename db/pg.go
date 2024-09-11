package db

import (
	"fmt"
	"log"
	"os"

	"github.com/JoiZs/go_todoapi/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnect() *gorm.DB {
	dbDSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DBhost"), os.Getenv("DBuser"), os.Getenv("DBpassword"), os.Getenv("DBdbname"), os.Getenv("DBport"))
	db, err := gorm.Open(postgres.Open(dbDSN))
	if err != nil {
		log.Fatal("Error at connecting db")
	}

	log.Println("Connected to DB")

	db.AutoMigrate(&model.Todo{})

	return db
}
