package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/JamesGopsill/swapiql/internal/resolvers"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	// Setting the logging level
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	debug := os.Getenv("GQL_DEBUG")
	if debug == "true" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// Booting up
	log.Info().Msg("SwapiQL booting up")

	// Read in the schema
	var s string
	files, err := ioutil.ReadDir("./schema")
	if err != nil {
		log.Fatal().Err(err).Msg("Schema Dir Error")
	}
	for _, file := range files {
		content, err := ioutil.ReadFile("./schema/" + file.Name())
		if err != nil {
			log.Fatal().Err(err).Msg("Schema File Error")
		}
		s += string(content)
	}

	// Add the options to the GraphQL schema
	opts := []graphql.SchemaOpt{
		graphql.UseFieldResolvers(),
		graphql.UseStringDescriptions(),
	}

	schema := graphql.MustParseSchema(s, &resolvers.RootResolver{
		BaseURL: "https://swapi.dev/api",
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

	log.Info().Msg("Listening on :3000 ...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal().Err(err)
	}

}
