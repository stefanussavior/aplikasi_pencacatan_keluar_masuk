package database

import (
	"aplikasi_pencatatan/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectionDatabase() {
	address := "root@tcp(127.0.0.1:3306)/aplikasi_pencatatan?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(address), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Data Terkoneksi dan table sudah termigrasi")
	DB.AutoMigrate(&models.User{}, &models.Pemasukan{}, &models.Pengeluaran{})
}
