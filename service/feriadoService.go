package service

import (
	"api-rest/types"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type FeriadosService struct {
	url string
}

func InitFeriadosService(url string) Service {
	return &FeriadosService{
		url: url,
	}
}

func (s *FeriadosService) GetFeriadosByAno(anio string) ([]types.Feriado, error) {
	resp, err := http.Get(s.url + "/" + anio)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var listaFeriados []types.Feriado
	body, err := io.ReadAll(resp.Body)
	json.Unmarshal([]byte(body), &listaFeriados)

	return listaFeriados, nil

}

func (*FeriadosService) GetNextFeriado(lista []types.Feriado) (*types.Feriado, error) {
	var nextFeriado types.Feriado
	for _, feriado := range lista {
		fecha, err := time.Parse("2006-01-02", feriado.Fecha)
		if err != nil {
			return nil, err
		}

		today := time.Now()
		if today.Before(fecha) || today.Equal(fecha) {
			nextFeriado = feriado
			break
		}
	}

	return &nextFeriado, nil

}
