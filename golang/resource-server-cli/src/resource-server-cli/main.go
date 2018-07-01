package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cloudfoundry-community/go-uaa"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	AirportsURL string `long:"airports-url" env:"AIRPORTS_URL" default:"https://localhost:9292"`
	UAAURL      string `long:"uaa-url" env:"UAA_URL" description:"Target UAA URL, e.g. https://login.mycompany.com:8443"`
	UAACACert   string `long:"uaa-ca-cert" env:"UAA_CA_CERT"`
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

	api, err := uaa.NewWithPasswordCredentials(opts.UAAURL, "", UAAClient, UAAClientSecret, opts.Username, opts.Password, uaa.JSONWebToken)
	if err != nil {
		log.Fatal(err)
	}
	api.Verbose = opts.Verbose
	if opts.UAACACert != "" {
		api.SkipSSLValidation = true
	}

	userinfo, err := api.GetMe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", userinfo)
}
