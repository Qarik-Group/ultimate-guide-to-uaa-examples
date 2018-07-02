package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"

	"github.com/cloudfoundry-community/go-uaa"
	"github.com/jessevdk/go-flags"
	"github.com/miquella/ask"
)

var opts struct {
	AirportsURL    string `long:"airports-url" env:"AIRPORTS_URL" default:"http://localhost:9292"`
	UAAURL         string `long:"uaa-url" env:"UAA_URL" description:"Target UAA URL, e.g. https://login.mycompany.com:8443"`
	UAACACert      string `long:"uaa-ca-cert" env:"UAA_CA_CERT"`
	UAAZoneID      string `long:"uaa-zone-id" env:"UAA_ZONE_ID"`
	Username       string `short:"u" long:"username" env:"UAA_USERNAME" description:"Username for authenticated user"`
	Password       string `short:"p" long:"password" env:"UAA_PASSWORD" description:"Password for authenticated user"`
	UserConfigPath string `long:"user-config" env:"AIRPORTS_USER_CONFIG" description:"Path to user auth config file" default:".user.json"`
	Verbose        bool   `short:"v" long:"verbose" description:"Show verbose debug information" env:"UAA_TRACE"`
}

type airport struct {
	AirportID          int     `json:"Airport ID"`
	Name               string  `json:"Name"`
	City               string  `json:"City"`
	Country            string  `json:"Country"`
	IATA               string  `json:"IATA"`
	ICAO               string  `json:"ICAO"`
	Latitude           float64 `json:"Latitude"`
	Longitude          float64 `json:"Longitude"`
	Altitude           int     `json:"Altitude"`
	Timezone           float64 `json:"Timezone"`
	DST                string  `json:"DST"`
	TzDatabaseTimeZone string  `json:"Tz database time zone"`
	Type               string  `json:"Type"`
	Source             string  `json:"Source"`
}

func main() {
	UAAClient := "airports"
	UAAClientSecret := "airports"

	parser := flags.NewParser(&opts, flags.Default)
	if _, err := parser.Parse(); err != nil {
		// log.Fatal(err)
		os.Exit(1)
	}

	if _, err := os.Stat(opts.UserConfigPath); os.IsNotExist(err) {
		username := opts.Username
		if username == "" {
			username, _ = ask.Ask("Username> ")
		}
		password := opts.Password
		if password == "" {
			password, _ = ask.HiddenAsk("Password> ")
		}

		creds, err := uaa.NewWithPasswordCredentials(opts.UAAURL, opts.UAAZoneID, UAAClient, UAAClientSecret, username, password, uaa.JSONWebToken)
		if err != nil {
			log.Fatal(err)
		}
		creds.API.Verbose = opts.Verbose
		if opts.UAACACert != "" {
			rootCAs, _ := x509.SystemCertPool()
			if rootCAs == nil {
				rootCAs = x509.NewCertPool()
			}
			rootCAs.AppendCertsFromPEM([]byte(opts.UAACACert))
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs: rootCAs,
				},
			}
			creds.API.UnauthenticatedClient = &http.Client{Transport: tr}
		}

		token, err := creds.Token()
		if err != nil {
			log.Fatal(err)
		}

		// To save to .user.json later
		tokenData, err := json.Marshal(token)
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile(opts.UserConfigPath, tokenData, 0600)
		if err != nil {
			log.Fatal(err)
		}
	}

	var token oauth2.Token
	storedToken, err := ioutil.ReadFile(opts.UserConfigPath)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(storedToken, &token)

	api, err := uaa.NewWithToken(opts.UAAURL, opts.UAAZoneID, token)
	api.Verbose = opts.Verbose
	if opts.UAACACert != "" {
		api.SkipSSLValidation = true
	}
	// me, err := api.GetMe()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%#v\n", me)

	// Call airports API with UAA authorization
	airportClient := &http.Client{Transport: http.DefaultTransport}
	req, err := http.NewRequest("GET", opts.AirportsURL, nil)
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", token.Type(), token.AccessToken))
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

	fmt.Printf("%#v\n", airports)
}
