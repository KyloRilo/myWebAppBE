package main

import(
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	// Schema
	fields := graphql.Fields{
		"test": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error){
				return "hello world", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	SchemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(SchemaConfig)
	if err != nil {
		log.Fatalf("Failed to create new schema, error: %v", err)
	}
	//Query
	query := `
		{
			test
		}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("Failed to execute graphql operation, errors: %v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON) //{"data":{"test":"hello world"}}

	//Handler
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}