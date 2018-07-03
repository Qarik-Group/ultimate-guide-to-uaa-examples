package actions

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"

	"github.com/gobuffalo/buffalo"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/cloudfoundry"
)

func init() {
	gothic.Store = App().SessionStore

	cf := cloudfoundry.New(os.Getenv("UAA_URL"), "airports-map", "airports-map",
		fmt.Sprintf("%s%s", App().Host, "/auth/cloudfoundry/callback"),
		"openid", "airports.all", "airports.50")

	if os.Getenv("UAA_CA_CERT") != "" {
		rootCAs, _ := x509.SystemCertPool()
		if rootCAs == nil {
			rootCAs = x509.NewCertPool()
		}
		rootCAs.AppendCertsFromPEM([]byte(os.Getenv("UAA_CA_CERT")))
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: rootCAs,
			},
		}
		cf.HTTPClient = &http.Client{Transport: tr}
	}

	goth.UseProviders(cf)
}

func Logout(c buffalo.Context) error {
	c.Session().Delete("email")
	c.Session().Delete("accessToken")
	c.Session().Delete("tokenType")
	return c.Redirect(302, "/")
}

func AuthCallback(c buffalo.Context) error {
	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Error(401, err)
	}
	c.Session().Set("email", user.Email)
	// c.Session().Set("scopes", user.Email) // TODO: user.Scopes
	c.Session().Set("accessToken", user.AccessToken)
	c.Session().Set("tokenType", "bearer") // TODO: user.TokenType
	return c.Redirect(302, "/")
}
