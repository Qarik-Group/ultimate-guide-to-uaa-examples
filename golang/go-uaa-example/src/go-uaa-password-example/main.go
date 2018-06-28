package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cloudfoundry-community/go-uaa"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	URL          string `long:"url" env:"UAA_URL" description:"Target UAA URL, e.g. https://login.mycompany.com:8443"`
	Client       string `short:"c" long:"client" env:"UAA_CLIENT" description:"UAA client that supports password grant" default:"our_uaa_cli"`
	ClientSecret string `short:"s" long:"client-secret" env:"UAA_CLIENT_SECRET" description:"UAA client secret" default:"our_uaa_cli_secret"`
	CACert       string `long:"ca-cert" env:"UAA_CA_CERT"`
	Username     string `short:"u" long:"username" env:"UAA_USERNAME" description:"Username for authenticated user"`
	Password     string `short:"p" long:"password" env:"UAA_PASSWORD" description:"Password for authenticated user"`
	Verbose      bool   `short:"v" long:"verbose" description:"Show verbose debug information" env:"UAA_TRACE"`
}

func main() {
	parser := flags.NewParser(&opts, flags.Default)
	if _, err := parser.Parse(); err != nil {
		// log.Fatal(err)
		os.Exit(1)
	}

	api, err := uaa.NewWithPasswordCredentials(opts.URL, "", opts.Client, opts.ClientSecret, opts.Username, opts.Password, uaa.JSONWebToken)
	if err != nil {
		log.Fatal(err)
	}
	api.Verbose = opts.Verbose
	if opts.CACert != "" {
		api.SkipSSLValidation = true
	}

	userinfo, err := api.GetMe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", userinfo)
}
