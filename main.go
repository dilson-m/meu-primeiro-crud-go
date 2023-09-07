package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

// Rota principal ou home
func rotaPrincipal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Bem vindo ao meu primeiro Crud em GoLang\n")
}

// Listar livros da biblioteca
func listarLivros(w http.ResponseWriter, r *http.Request) {
	//	setar cconfiguracao para apresentar retorno do json amigavel
	w.Header().Set("content-type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.Encode(Livros)
}

// Inclcuir livro na biblioteca
func cadastrarLivro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	//	Monstra o 1o registro setado na posiccao 0
	// encoder := json.NewEncoder(w)
	// encoder.Encode(Livros[0])

	body, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		//	tratar erro aqui
	}

	var novoLivro Livro
	json.Unmarshal(body, &novoLivro)
	novoLivro.Id = len(Livros) + 1
	Livros = append(Livros, novoLivro)

	encoder := json.NewEncoder(w)
	encoder.Encode(novoLivro)
}

// Define rotas / operacoes da API
func rotearLivros(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		listarLivros(w, r)
	} else if r.Method == "POST" {
		// fmt.Println("Home")
		cadastrarLivro(w, r)
	}
}

// Congiracao de rotas
func configurarRotas() {
	http.HandleFunc("/", rotaPrincipal)
	http.HandleFunc("/livros", rotearLivros)
}

// Define e habilita sesvidor HTTP
func configurarServidor() {
	configurarRotas()

	fmt.Printf("Servidor esta rodando na porta: %v \n", port)

	//	lod.fatal retorna erro caso nao consiga subir servidor http
	log.Fatal(http.ListenAndServe(":1337", nil)) //	DefaultServerMux
}

// Funcao principal
func main() {
	configurarServidor()
}
