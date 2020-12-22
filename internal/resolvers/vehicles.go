package resolvers

import (
	"github.com/rs/zerolog/log"
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
}

// ID creates the ID field equivalent to the number in parsed to SWAPI
func (v *VehicleResolver) ID() int32 {
	return GetIDFromURL(v.URL)
}

// Films resolves the films for a starship
func (v *VehicleResolver) Films() (*[]*FilmResolver, error) {
	return GetFilm(v.FilmURLs)
}

// Pilots resolves the pilots who flew the starship
func (v *VehicleResolver) Pilots() (*[]*PersonResolver, error) {
	return GetPerson(v.PilotURLs)
}

// GetVehicle requests a vehicle from the REST API
func GetVehicle(urls []string) (*[]*VehicleResolver, error) {

	var resolvers []*VehicleResolver
	var err error

	for _, url := range urls {
		var resolver VehicleResolver

		log.Debug().Str("URL", url).Msg("GetVehicle")

		err := GetURL(url, &resolver)
		if err != nil {
			return nil, err
		}

		resolvers = append(resolvers, &resolver)
	}

	return &resolvers, err

}
