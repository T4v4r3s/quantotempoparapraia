package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/cookies"
)

// Logger escreve informações da requisição no terminal
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		// Aqui ficam as instruções que serão executadas antes da requisição chegar ao controller
		// ou seja, é um middleware
		proximaFuncao(w, r)
		// Aqui ficam as instruções que serão executadas depois da requisição chegar ao controller
		// ou seja, é um middleware
	}
}

// Verificar se os dados estão no cookie, e não se estão corretos (API que faz isso)
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, erro := cookies.Ler(r); erro != nil {
			cookies.Deleta(w)
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
			fmt.Println(erro)
			return
		}

		//fmt.Println(valores, erro)

		// Aqui ficam as instruções que serão executadas antes da requisição chegar ao controller
		// ou seja, é um middleware
		proximaFuncao(w, r)
		// Aqui ficam as instruções que serão executadas depois da requisição chegar ao controller
		// ou seja, é um middleware
	}
}
