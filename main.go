package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/JamesGopsill/swapiql/internal/resolvers"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {
	// Booting up
	log.Println("Booting up swapiql")

	// Read in the schema
	var s string
	files, err := ioutil.ReadDir("./schema")
	if err != nil {
		log.Fatalf("Schema Dir Error: %v", err)
	}
	for _, file := range files {
		content, err := ioutil.ReadFile("./schema/" + file.Name())
		if err != nil {
			log.Fatalf("Schema File Error: %v", err)
		}
		s += string(content)
	}

	// Create the http client
	client := http.Client{Timeout: time.Second * 30}
	opts := []graphql.SchemaOpt{
		graphql.UseFieldResolvers(),
		graphql.UseStringDescriptions(),
	}

	schema := graphql.MustParseSchema(s, &resolvers.RootResolver{
		BaseURL: "https://swapi.dev/api",
		Client:  &client,
	}, opts...)

	//examples.ExampleHello(schema)
	//examples.ExamplePersonQuery(schema)
	//examples.ExampleFilmQuery(schema)
	//examples.ExamplePlanetQuery(schema)
	//examples.ExampleStarshipQuery(schema)
	//examples.ExampleVehicleQuery(schema)
	//examples.ExampleSpeciesQuery(schema)

	Server(schema)

}

// Server creates a DevServer
func Server(schema *graphql.Schema) {

	// Get the prefix from the Envvar if there is one
	prefix := os.Getenv("GQL_PREFIX")

	http.Handle(prefix+"/", &relay.Handler{Schema: schema})
	http.HandleFunc(prefix+"/playground", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/playground.html")
	})
	http.HandleFunc(prefix+"/voyager", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/voyager.html")
	})

	log.Println("Listening on :3000 ...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
