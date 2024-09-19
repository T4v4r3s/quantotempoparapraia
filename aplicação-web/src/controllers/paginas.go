package controllers

import (
	"net/http"
	"webapp/src/utils"
)

//Aqui ficam todas as funções que carregam páginas!

// CarregarPaginaPrincipal carrega a página principal
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {

	utils.ExecutarTemplate(w, "home.html", nil) //Carrega a página de login não passando dados para ela!

}
