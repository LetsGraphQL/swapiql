package resolvers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
)

// SpeciesResolver ...
type SpeciesResolver struct {
	Name            string   `json:"name"`
	URL             string   `json:"url"`
	Classification  string   `json:"classification"`
	Designation     string   `json:"designation"`
	AverageHeight   string   `json:"average_height"`
	AverageLifespan string   `json:"average_lifespan"`
	EyeColors       string   `json:"eye_colors"`
	HairColors      string   `json:"hair_colors"`
	SkinColors      string   `json:"skin_colors"`
	Language        string   `json:"language"`
	HomeworldURL    string   `json:"homeworld"`
	PeopleURLs      []string `json:"people"`
	FilmURLs        []string `json:"films"`
	Created         string   `json:"created"`
	Edited          string   `json:"edited"`

	Client *http.Client `json:"-"`
}

// ID creates the ID field equivalent to the number in parsed to SWAPI
func (s *SpeciesResolver) ID() int32 {
	return GetIDFromURL(s.URL)
}

// Homeworld ...
func (s *SpeciesResolver) Homeworld() (*[]*PlanetResolver, error) {
	return GetPlanet(s.Client, []string{s.HomeworldURL})
}

// People ...
func (s *SpeciesResolver) People() (*[]*PersonResolver, error) {
	return GetPerson(s.Client, s.PeopleURLs)
}

// Films ...
func (s *SpeciesResolver) Films() (*[]*FilmResolver, error) {
	return GetFilm(s.Client, s.FilmURLs)
}

// GetSpecies requests a species from the REST API
func GetSpecies(c *http.Client, urls []string) (*[]*SpeciesResolver, error) {

	var resolvers []*SpeciesResolver
	var err error

	if c == nil {
		err = errors.New("No client detected in resolver")
		return nil, err
	}

	for _, url := range urls {
		var resolver SpeciesResolver
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
