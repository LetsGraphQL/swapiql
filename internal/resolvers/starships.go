package resolvers

import (
	gocache "github.com/patrickmn/go-cache"
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

// GetStarship requests a starship from the REST API or cache
func GetStarship(urls []string) (*[]*StarshipResolver, error) {

	var resolvers []*StarshipResolver
	var err error

	for _, url := range urls {

		// Check the cache
		val, found := cache.Get(url)

		if found {

			log.Debug().Str("URL", url).Msg("GetStarship: Using cache")

			resolver := val.(*StarshipResolver)
			resolvers = append(resolvers, resolver)

		} else {

			log.Debug().Str("URL", url).Msg("GetStarship: Using REST API")

			var resolver StarshipResolver

			err := GetURL(url, &resolver)
			if err != nil {
				return nil, err
			}

			cache.Set(url, &resolver, gocache.DefaultExpiration)
			resolvers = append(resolvers, &resolver)

		}
	}

	return &resolvers, err

}

// SearchStarship searches for starships from the REST API
func SearchStarship(url string) (*[]*StarshipResolver, error) {

	var err error
	var r []*StarshipResolver
	var result struct {
		SearchResponse
		Results []*StarshipResolver `json:"results"`
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
			Results []*StarshipResolver `json:"results"`
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
