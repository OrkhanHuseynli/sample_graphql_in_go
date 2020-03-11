package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"sample_graphql_in_go/src/controller"
	"sample_graphql_in_go/src/graphqlapi"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main() {

	//Query
	query := `
		{
			payment(author: "Mitchel") {
					author
				}
		}
	`

	gqlprocessor := graphqlapi.NewGraphQLProcessor()
	gqlprocessor.ExecuteQuery(query)

		db, err := gorm.Open("mysql", "user:password@(db:3306)/db?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			log.Println(err)
			panic("failed to connect database")
		}
		defer db.Close()

		// Migrate the schema
		db.AutoMigrate(&Product{})

		// Create
		db.Create(&Product{Code: "L1212", Price: 1000})

		// Read
		var product Product
		db.First(&product, 1) // find product with id 1
		db.First(&product, "code = ?", "L1212") // find product with code l1212

		// Update - update product's price to 2000
		db.Model(&product).Update("Price", 2000)

		// Delete - delete product
		db.Delete(&product)

		port := "5000"
		service := controller.NewService(port)
		service.Run()
}