package controllers

import (
	"html/template"
	"log"
	"net/http"
	"std/github.com/danfmsouza/lojaGo/models"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.FindAllProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func Novo(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.FindAllProdutos()
	temp.ExecuteTemplate(w, "Novo", todosOsProdutos)
}

func Gravar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		parsePreco, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Não foi possível converter preco")
		}
		parseQuantidade, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Não foi possível converter quantidade")
		}
		models.GravarProduto(nome, descricao, parsePreco, parseQuantidade)
	}
	http.Redirect(w, r, "/", 301)
}

func Remove(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.RemoveProduto(idProduto)
	http.Redirect(w, r, "/", 301)
}

func Alterar(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produto := models.AlteraProduto(idProduto)
	temp.ExecuteTemplate(w, "Alterar", produto)
}
func Atualizar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		parseId, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Não foi possível converter o ID")
		}
		parsePreco, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Não foi possível converter preco")
		}
		parseQuantidade, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Não foi possível converter quantidade")
		}
		models.AtualizarProduto(parseId, nome, descricao, parsePreco, parseQuantidade)
	}
	http.Redirect(w, r, "/", 301)
}
