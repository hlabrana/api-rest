package main

import (
	"log"
	"net/http"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Hello World!" }`))
}

func main() {

	svc := initFeriadosService("https://apis.digital.gob.cl/fl/feriados")

	apiServer := InitApiServer(svc)
	log.Fatal(apiServer.Start(":3030"))

}
