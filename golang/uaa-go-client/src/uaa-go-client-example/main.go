package main

import (
	"fmt"
	"log"
	"os"

	"code.cloudfoundry.org/clock"
	"code.cloudfoundry.org/lager"

	client "code.cloudfoundry.org/uaa-go-client"
	config "code.cloudfoundry.org/uaa-go-client/config"
)

func main() {
	cfg := &config.Config{
		ClientName:       "our_uaa_cli",
		ClientSecret:     "our_uaa_cli_secret",
		UaaEndpoint:      "https://192.168.50.6:8443",
		SkipVerification: true,
		CACerts:          "",
	}

	logger := lager.NewLogger("test")
	clock := clock.NewClock()

	uaaClient, err := client.NewClient(logger, cfg, clock)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Printf("Connecting to: %s ...\n", cfg.UaaEndpoint)

	token, err := uaaClient.FetchToken(true)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Printf("Token: %#v\n", token)
}
