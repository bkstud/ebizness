package google

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"goauth/handlers/auth"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"os"
)

var cred auth.Credentials
var conf *oauth2.Config

func init() {
	file, err := ioutil.ReadFile("./google-creds.json")
	if err != nil {
		log.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	if err := json.Unmarshal(file, &cred); err != nil {
		log.Println("unable to marshal data")
		return
	}

	conf = &oauth2.Config{
		ClientID:     cred.Cid,
		ClientSecret: cred.Csecret,
		RedirectURL:  "http://127.0.0.1:1323/auth/google",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email", // You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
		},
		Endpoint: google.Endpoint,
	}
}

// LoginHandler handles the login procedure.
func LoginHandler(c echo.Context) error {
	log.Println("google login handler")
	return auth.LoginHandler(c, conf)
}

func AuthHandler(c echo.Context) error {
	log.Printf("google auth handler")
	return auth.AuthHandler(c, conf)
}
