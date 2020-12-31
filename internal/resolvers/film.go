package resolvers

import (
	"github.com/rs/zerolog/log"
)

// FilmResolver resolves a film
type FilmResolver struct {
	Title         string   `json:"title"`
	Episode       int32    `json:"episode_id"`
	OpeningCrawl  string   `json:"opening_crawl"`
	Director      string   `json:"director"`
	ProducerCSV   string   `json:"producer"`
	ReleaseDate   string   `json:"release_date"`
	Created       string   `json:"created"`
	URL           string   `json:"url"`
	CharacterURLs []string `json:"characters"`
	PlanetURLs    []string `json:"planets"`
	SpeciesURLs   []string `json:"species"`
}

// ID creates the ID field equivalent to the number in parsed to SWAPI
func (f *FilmResolver) ID() int32 {
	return GetIDFromURL(f.URL)
}

// Characters resolves the people in the film
func (f *FilmResolver) Characters() (*[]*PersonResolver, error) {
	return GetPerson(f.CharacterURLs)
}

// Planets resolves the planets in the film
func (f *FilmResolver) Planets() (*[]*PlanetResolver, error) {
	return GetPlanet(f.PlanetURLs)
}

// Species resolves the species for a person
func (f *FilmResolver) Species() (*[]*SpeciesResolver, error) {
	return GetSpecies(f.SpeciesURLs)
}

// Producers resolves the producers in the film
func (f *FilmResolver) Producers() *[]string {
	return SplitAndTrim(f.ProducerCSV)
}

// GetFilm requests a film from the REST API
func GetFilm(urls []string) (*[]*FilmResolver, error) {

	var resolvers []*FilmResolver
	var err error

	for _, url := range urls {
		var resolver FilmResolver

		log.Debug().Str("URL", url).Msg("GetFilm")

		err := GetURL(url, &resolver)
		if err != nil {
			return nil, err
		}

		resolvers = append(resolvers, &resolver)
	}

	return &resolvers, err
}

// SearchFilm searches for films from the REST API
func SearchFilm(url string) (*[]*FilmResolver, error) {

	var err error
	var r []*FilmResolver
	var result struct {
		SearchResponse
		Results []*FilmResolver `json:"results"`
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
			Results []*FilmResolver `json:"results"`
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
