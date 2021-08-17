package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Serrote", Descricao: "Utilizado para Cortar coisas", Preco: 50.9, Quantidade: 1},
		{"Sacos de Lixo", "Capacidade 100L", 9.1, 50},
		{"Pá", "Resistente", 49.9, 1},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}
