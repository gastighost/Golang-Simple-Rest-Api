package configs

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "host=localhost user=gaston password=password dbname=book-management-system port=6543 sslmode=disable TimeZone=Asia/Shanghai"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDb() *gorm.DB {
	return db
}
