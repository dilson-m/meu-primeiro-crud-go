package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var port = 1337

type Livro struct {
	Id     int    `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
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
	encoder := json.NewEncoder(w)
	encoder.Encode(Livros)
}

// Inclcuir livro na biblioteca
func cadastrarLivro(w http.ResponseWriter, r *http.Request) {
	// alatera o status da requisicao de 200(StatusOK) para 201(StatusCreated)
	w.WriteHeader(http.StatusCreated)

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

func deleteLivro(w http.ResponseWriter, r *http.Request) {
	partes := strings.Split(r.URL.Path, "/")
	id, erro := strconv.Atoi(partes[2])

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	indiceDoLivro := -1
	for indice, livro := range Livros {
		if livro.Id == id {
			indiceDoLivro = indice
			break
		}
	}

	if indiceDoLivro < 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ladoEsquerdo := Livros[0:indiceDoLivro]
	ladoDireito := Livros[indiceDoLivro+1 : len(Livros)]

	Livros = append(ladoEsquerdo, ladoDireito...)

	w.WriteHeader(http.StatusNoContent)
}

// Define rotas / operacoes da API
func rotearLivros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	partes := strings.Split(r.URL.Path, "/")

	if len(partes) == 2 || len(partes) == 3 && partes[2] == "" {
		if r.Method == "GET" {
			listarLivros(w, r)
		} else if r.Method == "POST" {
			cadastrarLivro(w, r)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	} else if len(partes) == 3 || len(partes) == 4 && partes[3] == "" {
		if r.Method == "GET" {
			buscarLivro(w, r)
		} else if r.Method == "DELETE" {
			deleteLivro(w, r)
		}
	}

}

// pesquisar livro por Id e apresentar
func buscarLivro(w http.ResponseWriter, r *http.Request) {
	partes := strings.Split(r.URL.Path, "/")
	id, _ := strconv.Atoi(partes[2])

	for _, livro := range Livros {
		if livro.Id == id {
			json.NewEncoder(w).Encode(livro)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

// Congiracao de rotas
func configurarRotas() {
	http.HandleFunc("/", rotaPrincipal)
	http.HandleFunc("/livros", rotearLivros)

	// e.g. /livros/123
	// http.HandleFunc("/livros/", buscarLivro)
	http.HandleFunc("/livros/", rotearLivros)

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
