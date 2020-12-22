package resolvers

import (
	"github.com/rs/zerolog/log"
)

// StarshipResolver resolves a starship from SWAPI
type StarshipResolver struct {
	Name                 string   `json:"name"`
	Model                string   `json:"model"`
	StarshipClass        string   `json:"starship_class"`
	ManufacturersCSV     string   `json:"manufacturer"`
	CostInCredits        string   `json:"cost_in_credits"`
	Length               string   `json:"length"`
	Crew                 string   `json:"crew"`
	Passengers           string   `json:"passengers"`
	MaxAtmospheringSpeed string   `json:"max_atmosphering_speed"`
	HyperdriveRating     string   `json:"hyperdrive_rating"`
	MGLT                 string   `json:"MGLT"`
	CargoCapacity        string   `json:"cargo_capacity"`
	Consumables          string   `json:"consumables"`
	URL                  string   `json:"url"`
	Created              string   `json:"created"`
	Edited               string   `json:"edited"`
	FilmURLs             []string `json:"films"`
	PilotURLs            []string `json:"pilots"`
}

// ID creates the ID field equivalent to the number in parsed to SWAPI
func (s *StarshipResolver) ID() int32 {
	return GetIDFromURL(s.URL)
}

// Films resolves the films for a starship
func (s *StarshipResolver) Films() (*[]*FilmResolver, error) {
	return GetFilm(s.FilmURLs)
}

// Pilots resolves the pilots who flew the starship
func (s *StarshipResolver) Pilots() (*[]*PersonResolver, error) {
	return GetPerson(s.PilotURLs)
}

// Manufacturers resolves the producers in the film
func (s *StarshipResolver) Manufacturers() *[]string {
	return SplitAndTrim(s.ManufacturersCSV)
}

// GetStarship requests a starship from the REST API
func GetStarship(urls []string) (*[]*StarshipResolver, error) {

	var resolvers []*StarshipResolver
	var err error

	for _, url := range urls {
		var resolver StarshipResolver

		log.Debug().Str("URL", url).Msg("GetStarship")

		err := GetURL(url, &resolver)
		if err != nil {
			return nil, err
		}

		resolvers = append(resolvers, &resolver)
	}

	return &resolvers, err

}
