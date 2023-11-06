// Você recebe uma lista de caminhos, onde `caminhos[i] = [cidadeAi, cidadeBi]` significa que existe um caminho direto que
//vai de cidadeAi para cidadeBi. Retorne a cidade de destino, ou seja, a cidade sem nenhum caminho que saia dela.

package main

import "fmt"

func Destino(caminhos [][2]string) string {
	cidadesOrigem := make(map[string]bool)

	for _, caminho := range caminhos {
		cidadesOrigem[caminho[0]] = true
	}

	var cidadeDest string
	for _, caminho := range caminhos {
		if !cidadesOrigem[caminho[1]] {
			cidadeDest = caminho[1]
			break
		}
	}

	return cidadeDest
}

func main() {
	caminhos := [][2]string{
		{"A", "B"},
		{"B", "C"},
		{"C", "D"},
		{"D", "E"},
	}

	destino := Destino(caminhos)

	fmt.Println("Cidade de destino:", destino)
}
