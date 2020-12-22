package resolvers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
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

	Client *http.Client `json:"-"`
}

// ID creates the ID field equivalent to the number in parsed to SWAPI
func (p *PlanetResolver) ID() int32 {
	return GetIDFromURL(p.URL)
}

// Residents resolves the people on the planet
func (p *PlanetResolver) Residents() (*[]*PersonResolver, error) {
	return GetPerson(p.Client, p.ResidentURLs)
}

// Films resolves the films for a planet
func (p *PlanetResolver) Films() (*[]*FilmResolver, error) {
	return GetFilm(p.Client, p.FilmURLs)
}

// Terrain resolves the planets terrain
func (p *PlanetResolver) Terrain() *[]string {
	return SplitAndTrim(p.TerrainCSV)
}

// GetPlanet requests a person from the REST API
func GetPlanet(c *http.Client, urls []string) (*[]*PlanetResolver, error) {

	var resolvers []*PlanetResolver
	var err error

	if c == nil {
		err = errors.New("No client detected in resolver")
		return nil, err
	}

	for _, url := range urls {
		var resolver PlanetResolver
		resolver.Client = c

		// Make sure it is using a secure connection
		if !strings.Contains(url, "https") {
			url = strings.ReplaceAll(url, "http", "https")
		}

		log.Printf("GetPlanet: %v", url)

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
