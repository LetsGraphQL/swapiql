package resolvers

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

// RootResolver is our root resolver for the GraphQL endpoint.
type RootResolver struct {
	BaseURL string
}

// Hello resolves the Hello Query
func (r *RootResolver) Hello() string {
	return "World"
}

// Person resolve the People Query
func (r *RootResolver) Person(ctx context.Context, args struct{ ID int32 }) (*[]*PersonResolver, error) {

	log.Debug().Int("ID", int(args.ID)).Msg("Resolving Person")
	url := []string{fmt.Sprintf("%v/people/%v/", r.BaseURL, args.ID)}
	return GetPerson(url)

}

// Film resolves the film query
func (r *RootResolver) Film(ctx context.Context, args struct{ ID int32 }) (*[]*FilmResolver, error) {

	log.Debug().Int("ID", int(args.ID)).Msg("Resolving Film")
	url := []string{fmt.Sprintf("%v/films/%v/", r.BaseURL, args.ID)}
	return GetFilm(url)

}

// Planet resolves the planet query
func (r *RootResolver) Planet(ctx context.Context, args struct{ ID int32 }) (*[]*PlanetResolver, error) {

	log.Debug().Int("ID", int(args.ID)).Msg("Resolving Planet")
	url := []string{fmt.Sprintf("%v/planets/%v/", r.BaseURL, args.ID)}
	return GetPlanet(url)

}

// Starship resolves the starship query
func (r *RootResolver) Starship(ctx context.Context, args struct{ ID int32 }) (*[]*StarshipResolver, error) {

	log.Debug().Int("ID", int(args.ID)).Msg("Resolving Starship")
	url := []string{fmt.Sprintf("%v/starships/%v/", r.BaseURL, args.ID)}
	return GetStarship(url)

}

// Vehicle resolves the vehicle query
func (r *RootResolver) Vehicle(ctx context.Context, args struct{ ID int32 }) (*[]*VehicleResolver, error) {

	log.Debug().Int("ID", int(args.ID)).Msg("Resolving Vehicle")
	url := []string{fmt.Sprintf("%v/vehicles/%v/", r.BaseURL, args.ID)}
	return GetVehicle(url)

}

// Species resolves the species query
func (r *RootResolver) Species(ctx context.Context, args struct{ ID int32 }) (*[]*SpeciesResolver, error) {

	log.Debug().Int("ID", int(args.ID)).Msg("Resolving Species")
	url := []string{fmt.Sprintf("%v/species/%v/", r.BaseURL, args.ID)}
	return GetSpecies(url)

}
