package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/miroslavLalev/cloudsync/src/db"
)

func registerSpaceRoutes(router *mux.Router) {
	router.Path("/{id}").Methods("GET").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		db.Test(vars["id"])
		w.Write([]byte("asd"))
	})
}
