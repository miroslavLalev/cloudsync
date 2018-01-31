package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/miroslavLalev/cloudsync/src/db/errors"
	"github.com/miroslavLalev/cloudsync/src/management"
	"github.com/miroslavLalev/cloudsync/src/objects"
	"github.com/miroslavLalev/cloudsync/src/server/dto"

	"github.com/gorilla/mux"
)

func registerSpaceRoutes(router *mux.Router) {
	router.Path("").Methods("POST").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var user dto.UsernameDto
		parseBody(req.Body, &user)
		err := management.GetRegistry().GetSpaceController().Create(user.Username)
		if err != nil {
			log.Println("Error during creation of space: ", err)
			if _, ok := err.(errors.SpaceAlreadyExists); ok {
				res.WriteHeader(http.StatusNotFound)
				res.Write([]byte(err.Error()))
				return
			}

			res.WriteHeader(http.StatusInternalServerError)
			return
		}

		res.WriteHeader(http.StatusNoContent)
	})

	router.Path("/{id}").Methods("GET").HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		id := mux.Vars(req)["id"]
		space, err := management.GetRegistry().GetSpaceController().Get(id)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte(err.Error()))
			return
		}

		res.Write(encodeSpace(space))
	})
}

func encodeSpace(s *objects.Space) []byte {
	j, err := json.Marshal(dto.NewSpace(s))
	if err != nil {
		log.Println(err)
	}

	return j
}
