package main

import "sample_graphql_in_go/src/graphqlapi"

func main() {

	// Query
	query := `
		{
			payment(author: "Mitchel") {
					author
				}
		}
	`

	gqlprocessor := graphqlapi.NewGraphQLProcessor()
	gqlprocessor.ExecuteQuery(query)
}