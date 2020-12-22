package resolvers

import (
	"log"
	"strconv"
	"strings"
)

// GetIDFromURL gets the id from the URL of the resource
func GetIDFromURL(url string) int32 {
	els := strings.Split(url, "/")
	if len(els) > 3 {
		i, err := strconv.ParseInt(els[len(els)-2], 10, 32)
		if err != nil {
			log.Println("Conversion Error")
		}
		return int32(i)
	}
	return int32(0)
}

// SplitAndTrim splits a csv delimited string and trims the whitespace
func SplitAndTrim(s string) *[]string {
	slc := strings.Split(s, ",")
	for i := range slc {
		slc[i] = strings.TrimSpace(slc[i])
	}
	return &slc
}
