package main

import "fmt"

func main(){
	frase := "O Junior foi a escola"
	tamanho := len(frase)
	fmt.Printf("tamanho da frase:%d\n", tamanho)
	primeiraFrase := frase[0:8]
	fmt.Printf("Primeira palavra: %s\n", primeiraFrase)
	segundaPalavra := frase[9:12]
	fmt.Printf("Segunda palavra:%s\n", segundaPalavra)
	ultimaPalavra := frase [12:]
	fmt.Printf("Ultima palavra: %s\n", ultimaPalavra)
}