package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var RotaPaginaPrincipal = []Rota{
	{
		URI:                "/",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaPrincipal,
		RequerAutenticacao: false,
	},
	{
		URI:                "/home",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaPrincipal,
		RequerAutenticacao: false,
	},
}
