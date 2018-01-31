package routes

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"

	"github.com/gorilla/mux"
)

const (
	spacesRoot    = "/spaces"
	resourcesRoot = "/resources"
)

func Register(router *mux.Router) *mux.Router {
	registerSpaceRoutes(router.PathPrefix(spacesRoot).Subrouter())
	registerResourceRoutes(router.PathPrefix(resourcesRoot).Subrouter())
	return router
}

// obj is a reference to where the body will be parsed
func parseBody(r io.Reader, obj interface{}) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		log.Println("Error reading body", err)
	}

	err = json.Unmarshal(body, obj)
	if err != nil {
		log.Println("Error parsing body", err)
	}
}
