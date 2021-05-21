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
	//CreateTable()
	//InsertData()
	//SearchData()
	//SearchAll()
	//SearchWhere()
	//Update()
	//UpdateWhere()
	Delete()
}

func CreateTable() {
	db.Migrator().DropTable(&student{})   // drop table
	db.Migrator().AutoMigrate(&student{}) // automatically creates column and table
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

	err := db.Create(&s).Error // checking for any error while create
	if err != nil {
		fmt.Println(err)
	}

}

func SearchData() {
	var s student
	err := db.First(&s).Error // return the first record from the db
	if err != nil {
		log.Fatal(err)

	}

	fmt.Println(s)

}

func SearchAll() {

	var s []student
	err := db.Find(&s).Error
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(s)

}

func SearchWhere() {

	var s student
	name := "Raj"
	err := db.Where("name = ?", name).First(&s).Error // if we have two names with raj
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(s)
}

func Update() {
	var s student

	db.First(&s) // filling the values in s var

	s.Name = "Amit" // accessing the struct field

	db.Save(&s) // update data

}

func UpdateWhere() {

	tx := db.Model(&student{})
	tx = tx.Where("email = ?", "diwakar@email.com")
	err := tx.Update("name", "Diwakar").Error
	if err != nil {
		log.Fatal(err)
	}
}
func Delete() {

	var s student
	db.First(&s)
	//db.Delete(&s) // soft delete // only updates timestamp and data remain as it is
	db.Unscoped().Delete(&s) // perm delete
}
