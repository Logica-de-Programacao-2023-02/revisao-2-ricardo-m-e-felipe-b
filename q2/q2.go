//O torneio de programação do CEUB ocorrerá em breve. Neste ano, equipes de quatro pessoas estão autorizadas a participar.
//
//No UniCEUB, temos um grupo de participantes que inclui programadores e matemáticos. Gostaríamos de saber o número máximo
//de equipes que podem ser formadas, considerando as seguintes regras:
//
//- Cada equipe deve ser composta por exatamente quatro estudantes.
//- Equipes compostas apenas por quatro matemáticos ou apenas por quatro programadores não têm um bom desempenho,
//  portanto, decidiu-se não formar tais equipes.
//- Assim, cada equipe deve ter pelo menos um programador e pelo menos um matemático.
//
//Escreva um programa que receba como entrada uma lista de participantes e retorne o número máximo de equipes que podem
//ser formadas, respeitando as regras mencionadas.
//
//Cada pessoa só pode fazer parte de uma equipe.

package main

type Equipe struct {
	Nome  string
	Curso string
}

func Verificar(equipes []Equipe) int {
	Programadores := 0
	Matematicos := 0

	for _, equipe := range equipes {
		if equipe.Curso == "Programador" {
			Programadores++
		} else if equipe.Curso == "Matematico" {
			Matematicos++
		}
	}

	MaxEquipes := 0

	if Programadores > 0 && Matematicos > 0 {
		MaxEquipes++
		Programadores--
		Matematicos--
	}

	MaxEquipes += (Programadores / 4) + (Matematicos / 4)

	return MaxEquipes
}

func main() {
	equipes := []Equipe{
		{"Felipe", "Programador"},
		{"Joao", "Programador"},
		{"Leticia", "Programador"},
		{"Arthur", "Matematico"},
	}
	maximoEquipes := Verificar(equipes)
	println("Número máximo de equipes que podem ser formadas:", maximoEquipes)
}
