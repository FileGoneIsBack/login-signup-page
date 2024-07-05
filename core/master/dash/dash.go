package dash

import (
	"log"
	"login/core/models"
	"login/core/models/functions"
	"login/core/models/server"
	"login/core/sessions"
	"net/http"
)

func init() {
	Route.NewSub(server.NewRoute("/", func(w http.ResponseWriter, r *http.Request) {
		type Page struct {
			Name, Title, Vers, Sitename       string
			*sessions.Session
		}
		ok, user := sessions.IsLoggedIn(w, r)
        if !ok {
            http.Redirect(w, r, "/auth/", http.StatusTemporaryRedirect)
            return
        }
		log.Printf("Username dash: %s", user.Username)
		functions.Render(Page{
			Title:        	"signup",
			Vers:        	models.Config.Vers,
			Sitename:		models.Config.Name,
			Session: user,
		}, w, "dash", "dash.html")
	}))
}
