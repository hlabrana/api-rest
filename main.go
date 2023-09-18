package main

import (
	"api-rest/handle"
	"api-rest/service"
	"log"
)

func main() {

	svc := service.InitFeriadosService("https://apis.digital.gob.cl/fl/feriados")

	apiServer := handle.InitApiServer(svc)
	log.Fatal(apiServer.Start(":3030"))

}
