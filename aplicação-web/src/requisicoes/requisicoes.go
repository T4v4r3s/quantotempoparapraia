package requisicoes

import (
	"io"
	"net/http"
	"webapp/src/cookies"
)

// FazerRequisicaoComAutenticacao faz uma requisição para a API, passando o token de autenticação
func FazerRequisicaoComAutenticacao(r *http.Request, metodo, url string, dados io.Reader) (*http.Response, error) {
	cliente := &http.Client{}
	request, erro := http.NewRequest(metodo, url, dados)
	if erro != nil {
		return nil, erro
	}
	cookie, _ := cookies.Ler(r)

	request.Header.Add("Authorization", "Bearer "+cookie["token"])
	//request.Header.Add("Authorization", "Bearer "+r.Header.Get("Authorization"))

	response, erro := cliente.Do(request)
	if erro != nil {
		return nil, erro
	}

	return response, nil
}
