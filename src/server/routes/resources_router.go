package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/miroslavLalev/cloudsync/src/objects"
	"github.com/miroslavLalev/cloudsync/src/server/dto"

	"github.com/gorilla/mux"
	"github.com/miroslavLalev/cloudsync/src/management"
)

func registerResourceRoutes(router *mux.Router) {
	router.Path("/{id}").Methods("GET").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		id := mux.Vars(req)["id"]
		resource, err := management.GetRegistry().GetResourceController().Get(id)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte(err.Error()))
			return
		}

		res.Write(encodeResource(resource))
	})
}

func encodeResource(r *objects.Resource) []byte {
	j, err := json.Marshal(dto.NewResourceDetailed(r))
	if err != nil {
		log.Println(err)
	}

	return j
}
