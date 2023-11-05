//O cozinheiro Remy preparou uma refeição para si mesmo e, enquanto almoçava, decidiu assistir a um vídeo no RataTube. No
//entanto, ele tem um tempo limitado de 𝑡 segundos para o almoço. Ele pediu a sua ajuda para escolher o vídeo ideal.
//
//O RataTube possui um feed de 𝑛 vídeos, cada um representado por uma estrutura de vídeo contendo informações sobre sua
//duração em segundos e o nível de entretenimento. O feed é inicialmente aberto no primeiro vídeo, e Remy pode pular para
//o próximo vídeo em 1 segundo (caso exista). Ele pode pular vídeos quantas vezes desejar, inclusive não pular nenhum.
//
//Sua tarefa é auxiliar Remy a escolher um vídeo que ele possa abrir e assistir dentro do tempo disponível, 𝑡 segundos. Se
//houver vários vídeos que se encaixam nessa condição, ele deseja escolher o vídeo com o maior nível de entretenimento.
//Retorne qualquer vídeo apropriado ou exiba um erro caso não haja um vídeo adequado dentro do tempo disponível.

package main

import (
	"errors"
	"fmt"
)

type Video struct {
	Nome          string
	Duration      int
	Entertainment int
}

func EscolherVideo(time int, videos []Video) (Video, error) {
	var (
		maxEntertainmentVideo Video
	)
	for i := range videos {
		video := videos[i]
		if video.Duration <= time && video.Entertainment > maxEntertainmentVideo.Entertainment {
			maxEntertainmentVideo = video
		}
		time--
	}
	if maxEntertainmentVideo.Entertainment == 0 {
		return Video{}, errors.New("nenhum video encontrado")
	}
	return maxEntertainmentVideo, nil
}

func main() {
	var t int
	fmt.Println("Quanto tempo você possui de almoço(em segundos) ? ")
	fmt.Scanln(&t)

	videos := []Video{
		{"Vídeo 1", 120, 5},
		{"Vídeo 2", 90, 4},
		{"Vídeo 3", 180, 7},
		{"Vídeo 4", 60, 3},
	}

	videoEscolhido, err := EscolherVideo(t, videos)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("Vídeo escolhido: %s\n", videoEscolhido.Nome)
	}
}
