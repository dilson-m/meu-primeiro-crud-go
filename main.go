package main

import (
	"fmt"
	"net/http"
)

var port = 1337

func rotaPrincipal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "meu primeiro Crud")
}

func configurarRotas() {
	http.HandleFunc("/", rotaPrincipal)
}

func configurarServidor() {
	configurarRotas()

	fmt.Print("Servidor esta rodando na porta: ", port)
	http.ListenAndServe(":1337", nil) //	DefaultServerMux
}

func main() {
	configurarServidor()
}
