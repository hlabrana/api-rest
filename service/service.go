package service

import "api-rest/types"

type Service interface {
	GetFeriadosByAno(anio string) ([]types.Feriado, error)
	GetNextFeriado(lista []types.Feriado) (*types.Feriado, error)
}
