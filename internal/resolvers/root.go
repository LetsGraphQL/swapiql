package resolvers

import (
	"context"
	"fmt"
	"net/url"

	"github.com/rs/zerolog/log"
)

// RootResolver is our root resolver for the GraphQL endpoint.
type RootResolver struct {
	BaseURL string
}

// Info resolves the Information about the service
func (r *RootResolver) Info() *InfoResolver {

	log.Debug().Msg("Resolving Info")
	requestsServed += 1
	return &InfoResolver{}

}

// Person resolve the People Query
func (r *RootResolver) Person(ctx context.Context, args struct {
	ID     *int32
	Search *string
}) (*[]*PersonResolver, error) {

	requestsServed += 1

	if args.ID != nil {

		id := *args.ID
		urls := []string{fmt.Sprintf("%v/people/%v/", r.BaseURL, id)}
		log.Debug().Str("URL", urls[0]).Msg("Resolving Person")
		return GetPerson(urls)

	}

	search := ""

	if args.Search != nil {
		search = *args.Search
	}

	urls := fmt.Sprintf("%v/people/?search=%v", r.BaseURL, url.QueryEscape(search))
	log.Debug().Str("Search Term", urls).Msg("Searching Person")
	return SearchPerson(urls)

}

// Film resolves the film query
func (r *RootResolver) Film(ctx context.Context, args struct {
	ID     *int32
	Search *string
}) (*[]*FilmResolver, error) {

	requestsServed += 1

	if args.ID != nil {

		id := *args.ID
		urls := []string{fmt.Sprintf("%v/films/%v/", r.BaseURL, id)}
		log.Debug().Str("URL", urls[0]).Msg("Resolving Films")
		return GetFilm(urls)

	}

	search := ""

	if args.Search != nil {
		search = *args.Search
	}

	urls := fmt.Sprintf("%v/films/?search=%v", r.BaseURL, url.QueryEscape(search))
	log.Debug().Str("Search Term", urls).Msg("Searching Films")
	return SearchFilm(urls)

}

// Planet resolves the planet query
func (r *RootResolver) Planet(ctx context.Context, args struct {
	ID     *int32
	Search *string
}) (*[]*PlanetResolver, error) {

	requestsServed += 1

	if args.ID != nil {

		id := *args.ID
		urls := []string{fmt.Sprintf("%v/planets/%v/", r.BaseURL, id)}
		log.Debug().Str("URL", urls[0]).Msg("Resolving Planets")
		return GetPlanet(urls)

	}

	search := ""

	if args.Search != nil {
		search = *args.Search
	}

	urls := fmt.Sprintf("%v/planets/?search=%v", r.BaseURL, url.QueryEscape(search))
	log.Debug().Str("Search Term", urls).Msg("Searching Planets")
	return SearchPlanet(urls)

}

// Starship resolves the starship query
func (r *RootResolver) Starship(ctx context.Context, args struct {
	ID     *int32
	Search *string
}) (*[]*StarshipResolver, error) {

	requestsServed += 1

	if args.ID != nil {

		id := *args.ID
		urls := []string{fmt.Sprintf("%v/starships/%v/", r.BaseURL, id)}
		log.Debug().Str("URL", urls[0]).Msg("Resolving Starships")
		return GetStarship(urls)

	}

	search := ""

	if args.Search != nil {
		search = *args.Search
	}

	urls := fmt.Sprintf("%v/starships/?search=%v", r.BaseURL, url.QueryEscape(search))
	log.Debug().Str("Search Term", urls).Msg("Searching Starships")
	return SearchStarship(urls)

}

// Vehicle resolves the vehicle query
func (r *RootResolver) Vehicle(ctx context.Context, args struct {
	ID     *int32
	Search *string
}) (*[]*VehicleResolver, error) {

	requestsServed += 1

	if args.ID != nil {

		id := *args.ID
		urls := []string{fmt.Sprintf("%v/vehicles/%v/", r.BaseURL, id)}
		log.Debug().Str("URL", urls[0]).Msg("Resolving Vehicles")
		return GetVehicle(urls)

	}

	search := ""

	if args.Search != nil {
		search = *args.Search
	}

	urls := fmt.Sprintf("%v/vehicles/?search=%v", r.BaseURL, url.QueryEscape(search))
	log.Debug().Str("Search Term", urls).Msg("Searching Vehicles")
	return SearchVehicle(urls)

}

// Species resolves the species query
func (r *RootResolver) Species(ctx context.Context, args struct {
	ID     *int32
	Search *string
}) (*[]*SpeciesResolver, error) {

	requestsServed += 1

	if args.ID != nil {

		id := *args.ID
		urls := []string{fmt.Sprintf("%v/species/%v/", r.BaseURL, id)}
		log.Debug().Str("URL", urls[0]).Msg("Resolving Species")
		return GetSpecies(urls)

	}

	search := ""

	if args.Search != nil {
		search = *args.Search
	}

	urls := fmt.Sprintf("%v/species/?search=%v", r.BaseURL, url.QueryEscape(search))
	log.Debug().Str("Search Term", urls).Msg("Searching Species")
	return SearchSpecies(urls)

}
