package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	product := NewProduct("Notebook", 1899.90)
	err = InsertProduct(db, product)

	if err != nil {
		panic(err)
	}

	product.Price = 100.0

	err = UpdateProduct(db, product)
	if err != nil {
		panic(err)
	}
}

func InsertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("INSERT INTO products(id, name, price) VALUES (?, ?, ?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	// res, err := stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}

	// res.LastInsertId()

	return nil
}

func UpdateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("UPDATE products SET name = ?, price = ? where id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}

	return nil
}
