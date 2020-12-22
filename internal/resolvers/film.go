package resolvers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
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

	Client *http.Client `json:"-"`
}

// ID creates the ID field equivalent to the number in parsed to SWAPI
func (f *FilmResolver) ID() int32 {
	return GetIDFromURL(f.URL)
}

// Characters resolves the people in the film
func (f *FilmResolver) Characters() (*[]*PersonResolver, error) {
	return GetPerson(f.Client, f.CharacterURLs)
}

// Planets resolves the planets in the film
func (f *FilmResolver) Planets() (*[]*PlanetResolver, error) {
	return GetPlanet(f.Client, f.PlanetURLs)
}

// Species resolves the species for a person
func (f *FilmResolver) Species() (*[]*SpeciesResolver, error) {
	return GetSpecies(f.Client, f.SpeciesURLs)
}

// Producers resolves the producers in the film
func (f *FilmResolver) Producers() *[]string {
	return SplitAndTrim(f.ProducerCSV)
}

// GetFilm requests a film from the REST API
func GetFilm(c *http.Client, urls []string) (*[]*FilmResolver, error) {

	var resolvers []*FilmResolver
	var err error

	if c == nil {
		err = errors.New("No client detected in resolver")
		return nil, err
	}

	for _, url := range urls {
		var resolver FilmResolver
		resolver.Client = c

		// Make sure it is using a secure connection
		if !strings.Contains(url, "https") {
			url = strings.ReplaceAll(url, "http", "https")
		}

		log.Printf("GetPerson: %v", url)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		res, err := c.Do(req)
		if err != nil {
			return nil, err
		}

		defer res.Body.Close()

		if res.StatusCode != 200 {
			err = errors.New(res.Status)
			return nil, err
		}

		err = json.NewDecoder(res.Body).Decode(&resolver)
		if err != nil {
			return nil, err
		}

		resolvers = append(resolvers, &resolver)

	}

	return &resolvers, err
}
