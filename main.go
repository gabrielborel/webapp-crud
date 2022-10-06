package main

import (
	"net/http"
	"webapp/routes"

	_ "github.com/lib/pq"
)


func main() {
	routes.LoadRoutes()

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic("Houve um erro ao tentar subir o servidor" + err.Error())
	}
}
