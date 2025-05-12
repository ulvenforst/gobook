package exercises

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

// The JSON-based web service of the Open Movie Database lets you search
// https://omdbapi.com/ for a movie by name and download its poster image.
// Write a tool poster that downloads the poster image for the movie named on the command line.
var posterURL struct{ Poster string }

// Poster fetches an image of a movie and downloads it into the root directory
func Poster(posterName string) error {

	// Constructs a valid url for our movie image
	fullURL, err := constructOMDbURL(posterName)
	if err != nil {
		return err
	}

	// We fetch the constructed URL and unmarchall the wanted value
	if err := json.Unmarshal(fetchURL(fullURL.String()), &posterURL); err != nil {
		return err
	}

	// Fetch the url of the image given from our initial URL
	posterIMG := fetchURL(posterURL.Poster)

	// We create the file to save the image data and save it
	fileName := posterName + ".jpg"
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	_, err = file.Write(posterIMG)
	if err != nil {
		return err
	}

	fmt.Printf("Poster image for %s saved as %s\n", posterName, fileName)
	return nil
}

func constructOMDbURL(target string) (*url.URL, error) {
	const BaseURL = "http://www.omdbapi.com/"
	params := url.Values{}
	params.Set("apikey", os.Getenv("OMDB_API_KEY"))
	params.Set("t", target)

	fullURL, err := url.Parse(BaseURL)
	if err != nil {
		return nil, err
	}
	fullURL.RawQuery = params.Encode()

	return fullURL, nil
}

func fetchURL(fullURL string) []byte {
	resp, err := http.Get(fullURL)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	return data
}
