package routes

import (
	"net/http"
	"std/github.com/danfmsouza/lojaGo/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/novo", controllers.Novo)
	http.HandleFunc("/gravar", controllers.Gravar)
	http.HandleFunc("/remover", controllers.Remove)
	http.HandleFunc("/alterar", controllers.Alterar)
	http.HandleFunc("/atualizar", controllers.Atualizar)

}
