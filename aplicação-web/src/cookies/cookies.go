package cookies

import (
	"net/http"
	"time"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// Configurar - Inicializa o pacote de cookies
func Configurar() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

// Salvar - Salva os cookies na resposta
func Salvar(w http.ResponseWriter, ID, token string) error {
	dados := map[string]string{
		"id":    ID,
		"token": token,
	}
	//dadosCodificados é um slice de bytes
	dadosCodificados, erro := s.Encode("dados", dados) //"dados" é o nome do cookie e dados é o valor do cookie
	if erro != nil {
		return erro
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    dadosCodificados,
		Path:     "/", //o cookie é válido em todas as rotas
		HttpOnly: true,
	}) //o cookie é salvo no navegador do usuário

	return nil
}

// Ler - Lê os cookies armazenados
func Ler(r *http.Request) (map[string]string, error) {
	cookie, erro := r.Cookie("dados")
	if erro != nil {
		return nil, erro
	}

	valores := make(map[string]string)
	if erro = s.Decode("dados", cookie.Value, &valores); erro != nil {
		return nil, erro
	}

	return valores, nil
}

// Deleta - Deleta os cookies salvos
func Deleta(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(-time.Hour),
		//MaxAge:   -1, //o cookie é deletado
	})
}

func VerificarEmBranco(r *http.Request) bool {
	cookie, err := Ler(r)
	if err != nil {
		return true
	}

	return cookie["token"] == ""
}
