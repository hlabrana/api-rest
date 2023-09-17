package main

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Service interface {
	GetFeriadosByAno(anio string) ([]Feriado, error)
	GetNextFeriado(lista []Feriado) (*Feriado, error)
}

type FeriadosService struct {
	url string
}

func initFeriadosService(url string) Service {
	return &FeriadosService{
		url: url,
	}
}

func (s *FeriadosService) GetFeriadosByAno(anio string) ([]Feriado, error) {
	resp, err := http.Get(s.url + "/" + anio)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var listaFeriados []Feriado
	body, err := io.ReadAll(resp.Body)
	json.Unmarshal([]byte(body), &listaFeriados)

	return listaFeriados, nil

}

func (*FeriadosService) GetNextFeriado(lista []Feriado) (*Feriado, error) {
	var nextFeriado Feriado
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
