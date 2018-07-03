package actions

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"code.cloudfoundry.org/clock"
	"code.cloudfoundry.org/lager"
	uaaclient "code.cloudfoundry.org/uaa-go-client"
	uaaconfig "code.cloudfoundry.org/uaa-go-client/config"
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

var uaa uaaclient.Client

func init() {
	uaaCfg := &uaaconfig.Config{
		ClientName:   "airports",
		ClientSecret: "airports",
		UaaEndpoint:  os.Getenv("UAA_URL"),
		CACerts:      os.Getenv("UAA_CA_CERT_FILE"),
	}
	logger := lager.NewLogger("airports")
	if os.Getenv("DEBUG") != "" {
		logger.RegisterSink(lager.NewPrettySink(os.Stdout, lager.DEBUG))
	}
	var err error
	uaa, err = uaaclient.NewClient(logger, uaaCfg, clock.NewClock())
	if err != nil {
		panic(err)
	}
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
	length := 10

	if c.Request().Header["Authorization"] != nil {
		authHeader := c.Request().Header["Authorization"][0]
		length = 20

		err := uaa.DecodeToken(authHeader, "openid")
		if err != nil {
			return err
		}
		if err := uaa.DecodeToken(authHeader, "airports.50"); err == nil {
			length = 50
		}
		if err := uaa.DecodeToken(authHeader, "airports.all"); err == nil {
			length = -1
		}
	}

	if length == -1 {
		return c.Render(200, r.JSON(airports))
	}
	return c.Render(200, r.JSON(airports[:length]))
}
