package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	db.Create(&Product{
		Name:  "Notebook",
		Price: 1000.00,
	})

	products := []Product{
		{
			Name:  "Notebook",
			Price: 1000.00,
		},
		{
			Name:  "Mouse",
			Price: 50.00,
		},
		{
			Name:  "Teclado",
			Price: 100.00,
		},
	}

	db.Create(&products)

	// var product Product
	// db.First(&product, 1)
	// fmt.Println(product)

	// db.First(&product, "name = ?", "Mouse")
	// fmt.Println(product)

	// var productss []Product
	// db.Find(&productss)

	// for _, pp := range productss {
	// 	fmt.Println(pp)
	// }

	// var productss []Product
	// db.Limit(2).Offset(2).Find(&productss)

	// for _, pp := range productss {
	// 	fmt.Println(pp)
	// }

	// var productss []Product
	// db.Where("price > ?", 100).Find(&productss)

	// for _, pp := range productss {
	// 	fmt.Println(pp)
	// }

	// var productss []Product
	// db.Where("name like ?", "%book%").Find(&productss)

	// for _, pp := range productss {
	// 	fmt.Println(pp)
	// }

	// var product Product
	// db.First(&product)

	// product.Name = "New Mouse"
	// db.Save(&product)

	// var product2 Product
	// db.First(&product2)

	// fmt.Println(product)
	// fmt.Println(product2)

	// db.Delete(&product2)

	var product Product
	db.First(&product)
	fmt.Println(product)
	product.Name = "New Mouse"
	db.Save(&product)
	db.Delete(&product)
}
