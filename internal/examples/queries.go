package examples

import (
	"context"
	"log"

	"github.com/graph-gophers/graphql-go"
)

// ExampleHello ...
func ExampleHello(schema *graphql.Schema) {

	queryString := "{ hello }"
	ExecQuery(queryString, schema)

}

// ExamplePersonQuery ...
func ExamplePersonQuery(schema *graphql.Schema) {

	queryString := "{ person(id: 1) { id, name, height, hairColor, skinColor, films { title }, homeworld { name } } }"
	ExecQuery(queryString, schema)

}

// ExampleFilmQuery ...
func ExampleFilmQuery(schema *graphql.Schema) {

	queryString := "{ film(id: 1) { id, title, episode, characters { name }, planets { name } } }"
	ExecQuery(queryString, schema)

}

// ExamplePlanetQuery ...
func ExamplePlanetQuery(schema *graphql.Schema) {

	queryString := "{ planet(id: 1) { name, diameter, residents { name }, films { title } } }"
	ExecQuery(queryString, schema)

}

// ExampleStarshipQuery ...
func ExampleStarshipQuery(schema *graphql.Schema) {

	queryString := "{ starship(id: 4) { name, films { title }, pilots { name } } }"
	ExecQuery(queryString, schema)

}

// ExampleVehicleQuery ...
func ExampleVehicleQuery(schema *graphql.Schema) {

	queryString := "{ vehicle(id: 4) { name } }"
	ExecQuery(queryString, schema)

}

// ExecQuery execute a query
func ExecQuery(query string, schema *graphql.Schema) {

	ctx := context.Background()
	var params map[string]interface{}
	r := schema.Exec(ctx, query, "", params)
	if r.Errors != nil {
		log.Println(r.Errors)
	}
	log.Println(string(r.Data))

}
