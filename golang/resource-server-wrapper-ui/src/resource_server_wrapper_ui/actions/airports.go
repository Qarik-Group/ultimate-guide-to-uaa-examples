package actions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

type airport struct {
	Name      string      `json:"Name"`
	City      string      `json:"City"`
	ICAO      string      `json:"ICAO"`
	Latitude  float64     `json:"Latitude"`
	Longitude float64     `json:"Longitude"`
	Altitude  int         `json:"Altitude"`
	Timezone  interface{} `json:"Timezone"`
	DST       string      `json:"DST"`
}

type FeatureCollection struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Type       string `json:"type"`
	Properties struct {
		Title string `json:"title"`
	} `json:"properties"`
	Geometry struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"geometry"`
}

// AirportFeatureCollectionHandler returns geojson FeatureCollection of airports
func AirportFeatureCollectionHandler(c buffalo.Context) error {
	airportsURL := "http://localhost:9292"
	airportClient := &http.Client{Transport: http.DefaultTransport}
	req, err := http.NewRequest("GET", airportsURL, nil)
	accessToken := c.Session().Get("accessToken")
	tokenType := c.Session().Get("tokenType")
	if accessToken != nil {
		req.Header.Set("Authorization", fmt.Sprintf("%s %s", tokenType, accessToken))
	}
	resp, err := airportClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var airports []airport
	err = json.Unmarshal(buf, &airports)
	if err != nil {
		log.Fatal(err)
	}

	// var features []Feature
	features := make([]Feature, len(airports))
	for i, airport := range airports {
		feature := &features[i]
		feature.Properties.Title = airport.ICAO
		feature.Geometry.Type = "Point"
		feature.Geometry.Coordinates = []float64{airport.Longitude, airport.Latitude}
	}
	collection := FeatureCollection{
		Type:     "FeatureCollection",
		Features: features,
	}
	return c.Render(200, r.JSON(collection))
}
