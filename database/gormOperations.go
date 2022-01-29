package gDB

import (
	"database/sql"
	"fmt"
	"keyserver/models"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var SqlDB *sql.DB

const StrNotFound = "norecords"

func Connect() bool {
	var err error
	var strDBConfig string

	log.Println("Establishing database connection")

	strDBConfig = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True", os.Getenv("DBUSER"), os.Getenv("DBPASS"),
		os.Getenv("DBHOST"), os.Getenv("DBNAME"))

	log.Println("DSN: " + strDBConfig)

	db, err = gorm.Open(mysql.Open(strDBConfig), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		return false
	}
	sqldb, err := db.DB()
	sqldb.SetMaxIdleConns(5)
	sqldb.SetMaxOpenConns(30)
	sqldb.SetConnMaxLifetime(time.Hour)

	log.Println("Connection established")
	return true
}

func Create(newEntity models.KeyData) error {
	return db.Model(&models.KeyData{}).Create(&newEntity).Error
}

func Get(keyData string) (models.KeyData, error) {
	var ret models.KeyData
	var count int64
	err := db.Model(&models.KeyData{}).Where(models.KeyData{Data: keyData}).First(&ret).Count(&count).Error
	if err != nil {
		return ret, err
	}
	if count > 0 {
		return ret, nil
	}
	return ret, fmt.Errorf(StrNotFound)
}
