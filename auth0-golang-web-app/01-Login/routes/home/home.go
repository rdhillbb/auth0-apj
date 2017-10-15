package home

import (
        templates "github.com/rdhillbb/auth0-apj/auth0-golang-web-app/01-Login/routes"
	"html/template"
	"net/http"
	"os"
	"fmt"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	data := struct {
		Auth0ClientId     string
		Auth0ClientSecret string
		Auth0Domain       string
		Auth0CallbackURL  template.URL
	}{
		os.Getenv("AUTH0_CLIENT_ID"),
		os.Getenv("AUTH0_CLIENT_SECRET"),
		os.Getenv("AUTH0_DOMAIN"),
		template.URL(os.Getenv("AUTH0_CALLBACK_URL")),
	}

	fmt.Println("Home----:",w)
	fmt.Println("Home----:rrrrrr",r)
	templates.RenderTemplate(w, "home", data)
}
