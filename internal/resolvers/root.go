package resolvers

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

// RootResolver is our root resolver for the GraphQL endpoint.
type RootResolver struct {
	BaseURL string
	Client  *http.Client
}

// Hello resolves the Hello Query
func (r *RootResolver) Hello() string {
	return "World"
}

// Person resolve the People Query
func (r *RootResolver) Person(ctx context.Context, args struct{ ID int32 }) (*[]*PersonResolver, error) {

	log.Println("Resolving Person with ID:", args.ID)
	url := []string{fmt.Sprintf("%v/people/%v/", r.BaseURL, args.ID)}
	return GetPerson(r.Client, url)

}

// Film resolves the film query
func (r *RootResolver) Film(ctx context.Context, args struct{ ID int32 }) (*[]*FilmResolver, error) {

	log.Println("Resolving Film with ID:", args.ID)
	url := []string{fmt.Sprintf("%v/films/%v/", r.BaseURL, args.ID)}
	return GetFilm(r.Client, url)

}

// Planet resolves the planet query
func (r *RootResolver) Planet(ctx context.Context, args struct{ ID int32 }) (*[]*PlanetResolver, error) {

	log.Println("Resolving Planet with ID:", args.ID)
	url := []string{fmt.Sprintf("%v/planets/%v/", r.BaseURL, args.ID)}
	return GetPlanet(r.Client, url)

}

// Starship resolves the starship query
func (r *RootResolver) Starship(ctx context.Context, args struct{ ID int32 }) (*[]*StarshipResolver, error) {

	log.Println("Resolving Starship with ID:", args.ID)
	url := []string{fmt.Sprintf("%v/starships/%v/", r.BaseURL, args.ID)}
	return GetStarship(r.Client, url)

}

// Vehicle resolves the vehicle query
func (r *RootResolver) Vehicle(ctx context.Context, args struct{ ID int32 }) (*[]*VehicleResolver, error) {

	log.Println("Resolving Vehicle with ID:", args.ID)
	url := []string{fmt.Sprintf("%v/vehicles/%v/", r.BaseURL, args.ID)}
	return GetVehicle(r.Client, url)

}

// Species resolves the species query
func (r *RootResolver) Species(ctx context.Context, args struct{ ID int32 }) (*[]*SpeciesResolver, error) {

	log.Println("Resolving Species with ID:", args.ID)
	url := []string{fmt.Sprintf("%v/species/%v/", r.BaseURL, args.ID)}
	return GetSpecies(r.Client, url)

}
