package main
//package main

import (
	"fmt"
	"time"
)

// Uma função simples que imprime algo com um atraso
func falar(texto string) {
	for i := 0; i < 3; i++ {
		fmt.Println(texto)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	// O comando 'go' antes da função cria uma Goroutine.
	// Isso faz a função rodar em "background" sem travar o programa!
	go falar("Aprendendo Go é incrível!")
	
	// Esta roda no fluxo principal
	falar("Go é muito rápido!")

	fmt.Println("Finalizado!")
}