package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type student struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique""`
}

var db *gorm.DB

func main() {

	newLogger := logger.New(
		log.New(os.Stdout, "\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			Colorful:      true,
			LogLevel:      logger.Info,
		},


	)
	var err error
	dsn := "host=localhost user=postgres password=root dbname=user port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})

	conn, _ := db.DB()
	defer conn.Close()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected")
	CreateTable()
	InsertData()

}

func CreateTable() {
	db.Migrator().DropTable(&student{})
	db.Migrator().AutoMigrate(&student{})
}

func InsertData() {
	s := []student{
		{
			Name:  "diwakar",
			Email: "diwakar@email.com",
		},
		{
			Name:  "Raj",
			Email: "raj@email.com",
		},
		{
			Name:  "dev",
			Email: "dev@email.com",
		},
	}

	err := db.Create(&s).Error // method in chaining
	if err != nil {
		fmt.Println(err)
	}

}
