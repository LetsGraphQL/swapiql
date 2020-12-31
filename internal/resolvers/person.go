package resolvers

import (
	"github.com/rs/zerolog/log"
)

// PersonResolver resolves a person from SWAPI
type PersonResolver struct {
	Name         string   `json:"name"`
	Height       string   `json:"height"`
	Mass         string   `json:"mass"`
	HairColor    string   `json:"hair_color"`
	SkinColor    string   `json:"skin_color"`
	EyeColor     string   `json:"eye_color"`
	BirthYear    string   `json:"birth_year"`
	Gender       string   `json:"gender"`
	Created      string   `json:"created"`
	URL          string   `json:"url"`
	FilmURLs     []string `json:"films"`
	StarshipURLs []string `json:"starships"`
	VehicleURLs  []string `json:"vehicles"`
	HomeworldURL string   `json:"homeworld"`
	SpeciesURLs  []string `json:"species"`
}

// ID creates the ID field equivalent to the number in parsed to SWAPI
func (p *PersonResolver) ID() int32 {
	return GetIDFromURL(p.URL)
}

// Films resolves the films for a person
func (p *PersonResolver) Films() (*[]*FilmResolver, error) {
	return GetFilm(p.FilmURLs)
}

// Homeworld resolves the planet for a person
func (p *PersonResolver) Homeworld() (*[]*PlanetResolver, error) {
	return GetPlanet([]string{p.HomeworldURL})
}

// Starships resolves the starships for a person
func (p *PersonResolver) Starships() (*[]*StarshipResolver, error) {
	return GetStarship(p.StarshipURLs)
}

// Vehicles resolves the vehicles for a person
func (p *PersonResolver) Vehicles() (*[]*VehicleResolver, error) {
	return GetVehicle(p.VehicleURLs)
}

// Species resolves the species for a person
func (p *PersonResolver) Species() (*[]*SpeciesResolver, error) {
	return GetSpecies(p.SpeciesURLs)
}

// GetPerson requests a person from the REST API
func GetPerson(urls []string) (*[]*PersonResolver, error) {

	var resolvers []*PersonResolver
	var err error

	for _, url := range urls {

		log.Debug().Str("URL", url).Msg("GetPerson")

		var resolver PersonResolver

		err := GetURL(url, &resolver)
		if err != nil {
			return nil, err
		}

		resolvers = append(resolvers, &resolver)
	}

	return &resolvers, err

}

// SearchPerson searches for people from the REST API
func SearchPerson(url string) (*[]*PersonResolver, error) {

	var err error
	var r []*PersonResolver
	var result struct {
		SearchResponse
		Results []*PersonResolver `json:"results"`
	}

	err = GetURL(url, &result)
	if err != nil {
		return nil, err
	}

	r = append(r, result.Results...)
	nextPage := result.Next
	log.Debug().Str("URL", nextPage).Msg("Next Page")

	// Loop if there is a next page
	for nextPage != "" {

		// Reset the struct
		var result struct {
			SearchResponse
			Results []*PersonResolver `json:"results"`
		}

		err = GetURL(nextPage, &result)
		if err != nil {
			return nil, err
		}

		r = append(r, result.Results...)
		nextPage = result.Next
		log.Debug().Str("URL", nextPage).Msg("Next Page")

	}

	return &r, err

}
