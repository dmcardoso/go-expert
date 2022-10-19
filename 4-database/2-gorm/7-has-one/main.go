package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	category := Category{Name: "Nova"}
	db.Create(&category)

	product := Product{
		Name:       "Mouse",
		Price:      1000.0,
		CategoryID: category.ID,
	}
	db.Create(&product)

	db.Create(&SerialNumber{
		Number:    "123456",
		ProductID: product.ID,
	})

	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)

	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	}

}
