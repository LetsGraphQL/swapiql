package resolvers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
)

// StarshipResolver resolves a starship from SWAPI
type StarshipResolver struct {
	Name                 string   `json:"name"`
	Model                string   `json:"model"`
	StarshipClass        string   `json:"starship_class"`
	Manufacturer         string   `json:"manufacturer"`
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

	Client *http.Client `json:"-"`
}

// ID creates the ID field equivalent to the number in parsed to SWAPI
func (s *StarshipResolver) ID() int32 {
	return GetIDFromURL(s.URL)
}

// Films resolves the films for a starship
func (s *StarshipResolver) Films() (*[]*FilmResolver, error) {
	return GetFilm(s.Client, s.FilmURLs)
}

// Pilots resolves the pilots who flew the starship
func (s *StarshipResolver) Pilots() (*[]*PersonResolver, error) {
	return GetPerson(s.Client, s.PilotURLs)
}

// GetStarship requests a starship from the REST API
func GetStarship(c *http.Client, urls []string) (*[]*StarshipResolver, error) {

	var resolvers []*StarshipResolver
	var err error

	if c == nil {
		err = errors.New("No client detected in resolver")
		return nil, err
	}

	for _, url := range urls {
		var resolver StarshipResolver
		resolver.Client = c

		// Make sure it is using a secure connection
		if !strings.Contains(url, "https") {
			url = strings.ReplaceAll(url, "http", "https")
		}

		log.Printf("GetVehicle: %v", url)

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
