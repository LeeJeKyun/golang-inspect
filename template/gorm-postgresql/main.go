package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=testdb sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Db 연결 실패")
	}
	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "123", Price: 1000})

	var product Product
	db.First(&product, 3)
	db.First(&product, "code = ?", "123")

	db.Model(&product).Update("price", 200)
	db.Model(&product).Updates(Product{Price: 250, Code: "456"})

	db.Delete(&product, 3)
}
