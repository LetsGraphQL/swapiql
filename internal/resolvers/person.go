package resolvers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
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

	Client *http.Client `json:"-"`
}

// ID creates the ID field equivalent to the number in parsed to SWAPI
func (p *PersonResolver) ID() int32 {
	return GetIDFromURL(p.URL)
}

// Films resolves the films for a person
func (p *PersonResolver) Films() (*[]*FilmResolver, error) {
	return GetFilm(p.Client, p.FilmURLs)
}

// Homeworld resolves the planet for a person
func (p *PersonResolver) Homeworld() (*[]*PlanetResolver, error) {
	return GetPlanet(p.Client, []string{p.HomeworldURL})
}

// Starships resolves the starships for a person
func (p *PersonResolver) Starships() (*[]*StarshipResolver, error) {
	return GetStarship(p.Client, p.StarshipURLs)
}

// Vehicles resolves the vehicles for a person
func (p *PersonResolver) Vehicles() (*[]*VehicleResolver, error) {
	return GetVehicle(p.Client, p.VehicleURLs)
}

// Species resolves the species for a person
func (p *PersonResolver) Species() (*[]*SpeciesResolver, error) {
	return GetSpecies(p.Client, p.SpeciesURLs)
}

// GetPerson requests a person from the REST API
func GetPerson(c *http.Client, urls []string) (*[]*PersonResolver, error) {

	var resolvers []*PersonResolver
	var err error

	if c == nil {
		err = errors.New("No client detected in resolver")
		return nil, err
	}

	for _, url := range urls {
		var resolver PersonResolver
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
