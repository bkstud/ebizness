package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

// Credentials which stores google ids.
type Credentials struct {
	Cid     string `json:"cid"`
	Csecret string `json:"csecret"`
}

func RandToken(l int) (string, error) {
	b := make([]byte, l)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func LoginHandler(c echo.Context, conf *oauth2.Config) error {
	log.Println("login handler")
	state, err := RandToken(32)
	if err != nil {
		return c.String(http.StatusInternalServerError, "error")
	}

	link := conf.AuthCodeURL(state)
	log.Printf("link:= %s", link)
	return c.Redirect(http.StatusTemporaryRedirect, link)
}

func AuthHandler(c echo.Context, conf *oauth2.Config) error {
	log.Printf("auth handler")
	code := c.QueryParam("code")
	log.Printf("code:= %v", code)
	if code == "" {
		log.Println("empty code")
		// Get token from database
	} else {
		tok, err := conf.Exchange(context.Background(), code)
		if err != nil {
			log.Println(err)
			out := map[string]string{"message": "Login failed. Please try again."}
			return c.JSON(http.StatusUnauthorized, out)
		}
		log.Printf("new user token:=%s", tok)
	}
	// conf.TokenSource(context.Background(), tok)
	// conf.clie
	return c.JSON(http.StatusOK, map[string]string{"message": "auth succeded"})
}
