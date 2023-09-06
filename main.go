package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var port = 1337

type Livro struct {
	Id     int
	Titulo string
	Autor  string
}

var Livros []Livro = []Livro{
	Livro{
		Id:     1,
		Titulo: "O Guarani",
		Autor:  "Jose de Alencar",
	},
	Livro{
		Id:     2,
		Titulo: "Cazuza",
		Autor:  "Viriato Correia",
	},
	Livro{
		Id:     3,
		Titulo: "Dom Casmurro",
		Autor:  "Machado de Assis",
	},
}

func rotaPrincipal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Bem vindo ao meu primeiro Crud em GoLang\n")
}

func listarLivros(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(Livros)
}

func configurarRotas() {
	http.HandleFunc("/Livros", listarLivros)
	http.HandleFunc("/", rotaPrincipal)
}

func configurarServidor() {
	configurarRotas()

	fmt.Printf("Servidor esta rodando na porta: %v \n", port)
	http.ListenAndServe(":1337", nil) //	DefaultServerMux
}

func main() {
	configurarServidor()
}
