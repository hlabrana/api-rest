package handle

import (
	"encoding/json"
	"net/http"

	"api-rest/service"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	svc service.Service
}

func InitApiServer(svc service.Service) *ApiServer {
	return &ApiServer{
		svc: svc,
	}
}

func (s *ApiServer) Start(listenAddr string) error {
	r := mux.NewRouter()
	r.HandleFunc("/feriados/{anio}", s.handleGetFeriadosByAnio)
	r.HandleFunc("/feriados-next", s.handleGetNextFeriado)
	return http.ListenAndServe(listenAddr, r)
}

func WriteJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
