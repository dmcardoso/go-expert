package main

import (
	"context"
	"database/sql"

	"github.com/dmcardoso/go-expert/15-sqlc/internal/db"
	"github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()

	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")

	if err != nil {
		panic(err)
	}

	defer dbConn.Close()

	queries := db.New(dbConn)

	id := uuid.New().String()
	err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:          id,
		Name:        "Backend",
		Description: sql.NullString{String: "Backend description", Valid: true},
	})

	if err != nil {
		panic(err)
	}

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          id,
		Name:        "Backend updated",
		Description: sql.NullString{String: "Backend description updated", Valid: true},
	})

	if err != nil {
		panic(err)
	}

	err = queries.DeleteCategory(ctx, id)

	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)

	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

}
