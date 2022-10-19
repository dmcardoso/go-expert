package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

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

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)

	defer cancel()

	selectProduct, err := SelectOneProduct(ctx, db, product.ID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Product(%v): %v is R$%.2f\n", selectProduct.ID, selectProduct.Name, selectProduct.Price)

	products, err := SelectAllProducts(ctx, db)
	if err != nil {
		panic(err)
	}

	for _, productSelected := range products {
		fmt.Printf("Product(%v): %v is R$%.2f\n", productSelected.ID, productSelected.Name, productSelected.Price)
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

func SelectOneProduct(ctx context.Context, db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var product Product

	// err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
	err = stmt.QueryRowContext(ctx, id).Scan(&product.ID, &product.Name, &product.Price)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func SelectAllProducts(ctx context.Context, db *sql.DB) ([]Product, error) {
	rows, err := db.QueryContext(ctx, "select id, name, price from products")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var product Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}
