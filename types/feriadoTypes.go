package types

type Feriados struct {
	Feriados []Feriado
}

type Feriado struct {
	Nombre        string `json:"nombre"`
	Comentarios   string `json:"comentarios"`
	Fecha         string `json:"fecha"`
	Irrenunciable string `json:"irrenunciable"`
	Tipo          string `json:"tipo"`
}
