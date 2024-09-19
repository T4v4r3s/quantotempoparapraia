package main

import (
	"fmt"
	"webapp/src/config"
	"webapp/src/utils"
)

//Cria uma hash e um block key para o cookie
/* func init() {
	hashkey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(hashkey)

	blockkey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(blockkey)
} */

func main() {
	config.Carregar()
	utils.CarregarTemplates() //pode ser feito numa função init também

	fmt.Printf("Rodando WebApp! Escutando na porta %d", config.Porta)

}
