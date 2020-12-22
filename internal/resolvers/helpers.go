package resolvers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Sets a global httpclient for the resolver package
var httpclient = &http.Client{Timeout: time.Second * 30}

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

// GetURL performs a GET request on a URL
func GetURL(url string, out interface{}) error {

	// Make sure it is using a secure connection
	if !strings.Contains(url, "https") {
		url = strings.ReplaceAll(url, "http", "https")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	res, err := httpclient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = errors.New(res.Status)
		return err
	}

	err = json.NewDecoder(res.Body).Decode(out)
	if err != nil {
		return err
	}

	return nil
}
