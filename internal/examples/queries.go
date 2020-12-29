package examples

import (
	"context"
	"log"

	"github.com/graph-gophers/graphql-go"
)

// ExamplePersonQuery performs a person query
func ExamplePersonQuery(schema *graphql.Schema) {

	queryString := "{ person(id: 1) { id, name, height, hairColor, skinColor, films { title }, homeworld { name } } }"
	ExecQuery(queryString, schema)

}

// ExampleFilmQuery performs a film query
func ExampleFilmQuery(schema *graphql.Schema) {

	queryString := "{ film(id: 1) { id, title, episode, characters { name }, planets { name } } }"
	ExecQuery(queryString, schema)

}

// ExamplePlanetQuery performs a planet query
func ExamplePlanetQuery(schema *graphql.Schema) {

	queryString := "{ planet(id: 1) { name, diameter, residents { name }, films { title } } }"
	ExecQuery(queryString, schema)

}

// ExampleStarshipQuery performs a starship query
func ExampleStarshipQuery(schema *graphql.Schema) {

	queryString := "{ starship(id: 4) { name, films { title }, pilots { name } } }"
	ExecQuery(queryString, schema)

}

// ExampleVehicleQuery performs a vehicle query
func ExampleVehicleQuery(schema *graphql.Schema) {

	queryString := "{ vehicle(id: 4) { name } }"
	ExecQuery(queryString, schema)

}

// ExampleSpeciesQuery performs a species query
func ExampleSpeciesQuery(schema *graphql.Schema) {

	queryString := "{ species(id: 4) { name, people { name } } }"
	ExecQuery(queryString, schema)

}

// ExampleInfoQuery performs a info query
func ExampleInfoQuery(schema *graphql.Schema) {

	queryString := "{ info() { title, uptime } }"
	ExecQuery(queryString, schema)

}

// ExecQuery executes a query
func ExecQuery(query string, schema *graphql.Schema) {

	ctx := context.Background()
	var params map[string]interface{}
	r := schema.Exec(ctx, query, "", params)
	if r.Errors != nil {
		log.Println(r.Errors)
	}
	log.Println(string(r.Data))

}
