//Na loja de animais à venda, existem alguns tipos de ração disponíveis para compra, sendo eles:
//
//* Ração para cachorro
//* Ração para gato
//* Ração universal
//
//O estoque dessas rações está representado em um mapa, onde a chave é o nome da ração e o valor correspondente é a
//quantidade disponível em estoque.
//
//Polycarp possui 𝑥 cães e 𝑦 gatos. Gostaríamos de determinar se é possível para ele comprar comida suficiente para todos
//os seus animais na loja. Cada um dos seus cães e gatos deve receber um pacote de ração adequado para sua espécie.

package main

import "fmt"

func Compra(dog, cat int, raçoes map[string]int) bool {

	if dog+cat <= raçoes["Ração Universal"] {
		return true
	}
	if dog > raçoes["Ração de cachorro"] {
		return false
	}
	if cat > raçoes["Ração de gato"] {
		return false
	}
	return true
}

func main() {

	raçoes := map[string]int{
		"Ração de cachorro": 4,
		"Ração de gato":     3,
		"Ração Universal":   8,
	}
	var cat int
	var dog int
	fmt.Println("Quantos gatos você possui ?")
	fmt.Scanln(&cat)
	fmt.Println("Quantos cachorros você possui ?")
	fmt.Scanln(&dog)

	if Compra(dog, cat, raçoes) {
		fmt.Println("As rações disponíveis na loja serão suficientes")
	} else {
		fmt.Println("Não há ração suficiente na loja para seus animais")
	}
}
