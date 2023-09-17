package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	svc Service
}

func InitApiServer(svc Service) *ApiServer {
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

func (s *ApiServer) handleGetFeriadosByAnio(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	anio, er := vars["anio"]
	if er != true {
		WriteJSON(w, http.StatusUnprocessableEntity, map[string]any{"error": "Falta ingresar Anio en URI"})
		return
	}

	res, err := s.svc.GetFeriadosByAno(anio)
	if err != nil {
		WriteJSON(w, http.StatusUnprocessableEntity, map[string]any{"error": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, res)
}

func (s *ApiServer) handleGetNextFeriado(w http.ResponseWriter, r *http.Request) {
	res, err := s.svc.GetFeriadosByAno(strconv.Itoa(time.Now().Year()))
	if err != nil {
		WriteJSON(w, http.StatusUnprocessableEntity, map[string]any{"error": err.Error()})
		return
	}

	feriado, err := s.svc.GetNextFeriado(res)
	if err != nil {
		WriteJSON(w, http.StatusUnprocessableEntity, map[string]any{"error": err.Error()})
		return
	}

	WriteJSON(w, http.StatusOK, feriado)

}

func WriteJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
