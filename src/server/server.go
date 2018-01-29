package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/miroslavLalev/cloudsync/src/server/routes"
)

const portVariable string = "PORT"

type RestServer struct {
	port string
}

func GetRestServer() *RestServer {
	return rs
}

var rs *RestServer

func (rs *RestServer) Start() {
	router := mux.NewRouter()
	routes.Register(router)

	http.Handle("/", router)
	err := http.ListenAndServe(":"+rs.port, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func init() {
	port := os.Getenv(portVariable)
	if port == "" {
		panic(fmt.Sprintf("Failed to read port from variable '%v'", portVariable))
	}

	rs = &RestServer{port: port}
}
