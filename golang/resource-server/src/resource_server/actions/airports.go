package actions

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/packr"
)

type airport struct {
	Name      string  `json:"Name"`
	City      string  `json:"City"`
	ICAO      string  `json:"ICAO"`
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
	Altitude  int     `json:"Altitude"`
}

// AirportsHandler returns a list of Australian airports
func AirportsHandler(c buffalo.Context) error {
	dataBox := packr.NewBox("../data")
	airportFile, err := dataBox.Open("airports.json")
	if err != nil {
		return err
	}
	airportJSON, err := ioutil.ReadAll(airportFile)
	if err != nil {
		return err
	}
	var airports []airport
	json.Unmarshal(airportJSON, &airports)
	return c.Render(200, r.JSON(airports))
}
