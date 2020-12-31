package resolvers

import (
	"github.com/rs/zerolog/log"
)

// PlanetResolver resolves a plane from SWAPI
type PlanetResolver struct {
	Name           string   `json:"name"`
	RotationPeriod string   `json:"rotation_period"`
	OrbitalPeriod  string   `json:"orbital_period"`
	Diameter       string   `json:"diameter"`
	Climate        string   `json:"climate"`
	Gravity        string   `json:"gravity"`
	TerrainCSV     string   `json:"terrain"`
	SurfaceWater   string   `json:"surface_water"`
	Population     string   `json:"population"`
	URL            string   `json:"url"`
	ResidentURLs   []string `json:"residents"`
	FilmURLs       []string `json:"films"`
}

// ID creates the ID field equivalent to the number in parsed to SWAPI
func (p *PlanetResolver) ID() int32 {
	return GetIDFromURL(p.URL)
}

// Residents resolves the people on the planet
func (p *PlanetResolver) Residents() (*[]*PersonResolver, error) {
	return GetPerson(p.ResidentURLs)
}

// Films resolves the films for a planet
func (p *PlanetResolver) Films() (*[]*FilmResolver, error) {
	return GetFilm(p.FilmURLs)
}

// Terrain resolves the planets terrain
func (p *PlanetResolver) Terrain() *[]string {
	return SplitAndTrim(p.TerrainCSV)
}

// GetPlanet requests a person from the REST API
func GetPlanet(urls []string) (*[]*PlanetResolver, error) {

	var resolvers []*PlanetResolver
	var err error

	for _, url := range urls {
		var resolver PlanetResolver

		log.Debug().Str("URL", url).Msg("GetPlanet")

		err := GetURL(url, &resolver)
		if err != nil {
			return nil, err
		}

		resolvers = append(resolvers, &resolver)
	}

	return &resolvers, err

}

// SearchPlanet searches for planets from the REST API
func SearchPlanet(url string) (*[]*PlanetResolver, error) {

	var err error
	var r []*PlanetResolver
	var result struct {
		SearchResponse
		Results []*PlanetResolver `json:"results"`
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
			Results []*PlanetResolver `json:"results"`
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

