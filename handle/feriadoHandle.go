package handle

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

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
