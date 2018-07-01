package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cloudfoundry-community/go-uaa"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	AirportsURL string `long:"airports-url" env:"AIRPORTS_URL" default:"https://localhost:9292"`
	UAAURL      string `long:"uaa-url" env:"UAA_URL" description:"Target UAA URL, e.g. https://login.mycompany.com:8443"`
	UAACACert   string `long:"uaa-ca-cert" env:"UAA_CA_CERT"`
	UAAZoneID   string `long:"uaa-zone-id" env:"UAA_ZONE_ID"`
	Username    string `short:"u" long:"username" env:"UAA_USERNAME" description:"Username for authenticated user"`
	Password    string `short:"p" long:"password" env:"UAA_PASSWORD" description:"Password for authenticated user"`
	Verbose     bool   `short:"v" long:"verbose" description:"Show verbose debug information" env:"UAA_TRACE"`
}

func main() {
	UAAClient := "airports"
	UAAClientSecret := "airports"

	parser := flags.NewParser(&opts, flags.Default)
	if _, err := parser.Parse(); err != nil {
		// log.Fatal(err)
		os.Exit(1)
	}

	creds, err := uaa.NewWithPasswordCredentials(opts.UAAURL, opts.UAAZoneID, UAAClient, UAAClientSecret, opts.Username, opts.Password, uaa.JSONWebToken)
	if err != nil {
		log.Fatal(err)
	}
	creds.API.Verbose = opts.Verbose
	if opts.UAACACert != "" {
		creds.API.SkipSSLValidation = true
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		creds.API.UnauthenticatedClient = &http.Client{Transport: tr}
	}

	token, err := creds.Token()
	if err != nil {
		log.Fatal(err)
	}

	// To save to .user.json later
	// s, err := json.Marshal(token)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(s))

	api, err := uaa.NewWithToken(opts.UAAURL, opts.UAAZoneID, *token)
	api.Verbose = opts.Verbose
	if opts.UAACACert != "" {
		api.SkipSSLValidation = true
	}
	me, err := api.GetMe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", me)
}
