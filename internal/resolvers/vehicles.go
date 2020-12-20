package resolvers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
)

// VehicleResolver resolves a vehicle from SWAPI
type VehicleResolver struct {
	Name                 string   `json:"name"`
	Model                string   `json:"model"`
	VehicleClass         string   `json:"vehicle_class"`
	Manufacturer         string   `json:"manufacturer"`
	Length               string   `json:"length"`
	CostInCredits        string   `json:"cost_in_credits"`
	Crew                 string   `json:"crew"`
	Passengers           string   `json:"passengers"`
	MaxAtmospheringSpeed string   `json:"max_atmosphering_speed"`
	CargoCapacity        string   `json:"capacity_crew"`
	Consumables          string   `json:"consumables"`
	Created              string   `json:"created"`
	Edited               string   `json:"edited"`
	URL                  string   `json:"url"`
	FilmURLs             []string `json:"films"`
	PilotURLs            []string `json:"pilots"`

	Client *http.Client `json:"-"`
}

// ID creates the ID field equivalent to the number in parsed to SWAPI
func (v *VehicleResolver) ID() int32 {
	return GetIDFromURL(v.URL)
}

// Films resolves the films for a starship
func (v *VehicleResolver) Films() (*[]*FilmResolver, error) {
	return GetFilm(v.Client, v.FilmURLs)
}

// Pilots resolves the pilots who flew the starship
func (v *VehicleResolver) Pilots() (*[]*PersonResolver, error) {
	return GetPerson(v.Client, v.PilotURLs)
}

// GetVehicle requests a vehicle from the REST API
func GetVehicle(c *http.Client, urls []string) (*[]*VehicleResolver, error) {

	var resolvers []*VehicleResolver
	var err error

	if c == nil {
		err = errors.New("No client detected in resolver")
		return nil, err
	}

	for _, url := range urls {
		var resolver VehicleResolver
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
