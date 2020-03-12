package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	port := "5000"
	service := controller.NewService(port)
	service.Run()
}