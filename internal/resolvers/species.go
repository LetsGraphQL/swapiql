package resolvers

import (
	"github.com/rs/zerolog/log"
)

// SpeciesResolver ...
type SpeciesResolver struct {
	Name            string   `json:"name"`
	URL             string   `json:"url"`
	Classification  string   `json:"classification"`
	Designation     string   `json:"designation"`
	AverageHeight   string   `json:"average_height"`
	AverageLifespan string   `json:"average_lifespan"`
	EyeColorsCSV    string   `json:"eye_colors"`
	HairColorsCSV   string   `json:"hair_colors"`
	SkinColorsCSV   string   `json:"skin_colors"`
	Language        string   `json:"language"`
	HomeworldURL    string   `json:"homeworld"`
	PeopleURLs      []string `json:"people"`
	FilmURLs        []string `json:"films"`
	Created         string   `json:"created"`
	Edited          string   `json:"edited"`
}

// ID creates the ID field equivalent to the number in parsed to SWAPI
func (s *SpeciesResolver) ID() int32 {
	return GetIDFromURL(s.URL)
}

// Homeworld ...
func (s *SpeciesResolver) Homeworld() (*[]*PlanetResolver, error) {
	return GetPlanet([]string{s.HomeworldURL})
}

// People ...
func (s *SpeciesResolver) People() (*[]*PersonResolver, error) {
	return GetPerson(s.PeopleURLs)
}

// Films ...
func (s *SpeciesResolver) Films() (*[]*FilmResolver, error) {
	return GetFilm(s.FilmURLs)
}

// EyeColors resolves the common eye colors for a species
func (s *SpeciesResolver) EyeColors() *[]string {
	return SplitAndTrim(s.EyeColorsCSV)
}

// SkinColors resolves the common skin colors for a species
func (s *SpeciesResolver) SkinColors() *[]string {
	return SplitAndTrim(s.SkinColorsCSV)
}

// HairColors resolves the common hair colors for a species
func (s *SpeciesResolver) HairColors() *[]string {
	return SplitAndTrim(s.HairColorsCSV)
}

// GetSpecies requests a species from the REST API
func GetSpecies(urls []string) (*[]*SpeciesResolver, error) {

	var resolvers []*SpeciesResolver
	var err error

	for _, url := range urls {
		var resolver SpeciesResolver

		log.Debug().Str("URL", url).Msg("GetSpecies")

		err := GetURL(url, &resolver)
		if err != nil {
			return nil, err
		}

		resolvers = append(resolvers, &resolver)
	}

	return &resolvers, err

}

// SearchSpecies searches for starships from the REST API
func SearchSpecies(url string) (*[]*SpeciesResolver, error) {

	var err error
	var r []*SpeciesResolver
	var result struct {
		SearchResponse
		Results []*SpeciesResolver `json:"results"`
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
			Results []*SpeciesResolver `json:"results"`
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
