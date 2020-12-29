package resolvers

import (
	"strconv"
	"time"
)

// InfoResolver resolves the information about the service
type InfoResolver struct {
}

// Title gets the title of the service
func (i *InfoResolver) Title() string {
	return "SWAPIQL"
}

// RepositoryURL gets the repository url of the service
func (i *InfoResolver) RepositoryURL() string {
	return "https://github.com/JamesGopsill/swapiql"
}

// DockerURL gets the docker url of the service
func (i *InfoResolver) DockerURL() string {
	return "https://hub.docker.com/r/jgopsill/swapiql"
}

// Description gets the docker url of the service
func (i *InfoResolver) Description() string {
	return "This repo is a proxy GraphQL wrapper service for the SWAPI REST API enabling the community to access their Star War trivia in GraphQL form."
}

// UpFrom returns the time the service started
func (i *InfoResolver) UpFrom() string {
	return upfrom.String()
}

// UpTime returns the time the service started
func (i *InfoResolver) UpTime() string {
	delta := time.Now().Sub(upfrom)
	return delta.String()
}

// RequestsServed returns the time the service started
func (i *InfoResolver) RequestsServed() string {
	return strconv.FormatInt(requestsServed, 10)
}
