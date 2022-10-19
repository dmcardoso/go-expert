package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
	gorm.Model
}

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
	gorm.Model
	Products []Product `gorm:"many2many:products_categories;"`
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	newCategory := Category{Name: "Nova"}
	db.Create(&newCategory)

	tx := db.Begin()
	var category Category

	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&category).Error

	if err != nil {
		panic(err)
	}

	category.Name = "Eletrônicos"

	tx.Debug().Save(&category)
	tx.Commit()

}
