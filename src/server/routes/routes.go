package routes

import (
	"github.com/gorilla/mux"
)

const (
	spacesRoot = "/spaces"
)

func Register(router *mux.Router) *mux.Router {
	registerSpaceRoutes(router.PathPrefix(spacesRoot).Subrouter())
	return router
}
