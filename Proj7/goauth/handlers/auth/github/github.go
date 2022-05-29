package github

import (
	"encoding/json"
	"goauth/handlers/auth"
	"io/ioutil"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var cred auth.Credentials
var conf *oauth2.Config

func init() {
	file, err := ioutil.ReadFile("./github-creds.json")
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
		RedirectURL:  "http://localhost:1323/auth/github",
		Endpoint:     github.Endpoint,
	}
}
func LoginHandler(c echo.Context) error {
	return auth.LoginHandler(c, conf)
}

func AuthHandler(c echo.Context) error {
	return auth.AuthHandler(c, conf)
}
