package main

/**
 * $ go run main.go
 * Please open the URL and re-run this command with the --auth-code flag to provide the auth code:
 * https://192.168.50.6:8443/oauth/authorize?client_id=uaa-cli-authcode&response_type=code
 *
 * $ go run main.go --auth-code PHIZTvnyp5
 * &uaa.UserInfo{UserID:"10661d47-6d08-4872-81aa-7868327cff0b", Sub:"10661d47-6d08-4872-81aa-7868327cff0b", Username:"drnic", GivenName:"Dr Nic", FamilyName:"Williams", Email:"drnic@starkandwayne", PhoneNumber:[]string(nil), PreviousLoginTime:1530165077339, Name:"Dr Nic Williams"}
 */
import (
	"fmt"
	"log"
	"os"

	"github.com/cloudfoundry-community/go-uaa"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	URL          string `long:"url" env:"UAA_URL" description:"Target UAA URL, e.g. https://login.mycompany.com:8443"`
	Client       string `short:"c" long:"client" env:"UAA_CLIENT" description:"UAA client that supports password grant" default:"uaa-cli-authcode"`
	ClientSecret string `short:"s" long:"client-secret" env:"UAA_CLIENT_SECRET" description:"UAA client secret" default:"uaa-cli-authcode"`
	CACert       string `long:"ca-cert" env:"UAA_CA_CERT"`
	AuthCode     string `long:"auth-code" description:"Authorization code returned from initial run"`
	Verbose      bool   `short:"v" long:"verbose" description:"Show verbose debug information" env:"UAA_TRACE"`
}

func main() {

	parser := flags.NewParser(&opts, flags.Default)
	if _, err := parser.Parse(); err != nil {
		// log.Fatal(err)
		os.Exit(1)
	}

	if opts.AuthCode == "" {
		authURL := fmt.Sprintf("%s/oauth/authorize?client_id=%s&response_type=code", opts.URL, opts.Client)
		fmt.Println("Please open the URL and re-run this command with the --auth-code flag to provide the auth code:")
		fmt.Println(authURL)
		os.Exit(0)
	}

	// Need to print something + redirect URL, and wait for return code
	code := opts.AuthCode
	api, err := uaa.NewWithAuthorizationCode(opts.URL, "", opts.Client, opts.ClientSecret, code, opts.CACert != "", uaa.JSONWebToken)
	if err != nil {
		log.Fatal(err)
	}
	api.Verbose = opts.Verbose

	userinfo, err := api.GetMe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", userinfo)
}
