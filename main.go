package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, "meu primeiro Crud")
	})
	http.ListenAndServe(":1337", nil) //	DefaultServerMux

}
