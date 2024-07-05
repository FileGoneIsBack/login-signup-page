package master

import (
	"login/core/master/authentication"
	"login/core/master/dash"
	"login/core/master/api"
	"login/core/master/landing"
	"login/core/models"
	"login/core/models/server"
	"net/http"
)

var (
	Service *server.Server = server.NewServer(&server.Config{
		Addr:   "0.0.0.0:80",
		Secure: models.Config.Secure,
		Cert:   models.Config.Cert,
		Key:    models.Config.Key,
	})
	Route  *server.Route   = server.NewSubRouter("")
	Assets *server.Handler = server.NewHandler("/_assets/", http.StripPrefix("/_assets", http.FileServer(http.Dir("assets/branding"))))
)

func NewV2() {
	Route.NewSubs(
		authentication.Route,
		dash.Route,
		landing.Route,
		api.Route,
	)
	Service.AddRoute(Route)
	Service.AddHandler(Assets)



	Service.AddRoute(server.NewRoute("", nil))
	if err := Service.ListenAndServe(); err != nil {
		panic(err)
	}
}