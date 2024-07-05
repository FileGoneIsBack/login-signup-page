package api

import (
	authentication "login/core/master/api/auth"
	"login/core/models/server"
)

var (
	Route *server.Route = server.NewSubRouter("/api")
)

func init() {
	Route.NewSubs(
		authentication.Route,
	)
}
