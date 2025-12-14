package main

import "fmt"

// %s = converte para string
// %c = converte para caractere

func main() {
	frase := "Ola', Mundo!!!"
	primeiroCaractere := frase[0]
	fmt.Printf("primeiro Caractere: %c\n", primeiroCaractere)

	indiceUltimoCaractere := len(frase) - 1
	ultimoCaracere := frase[indiceUltimoCaractere]
	fmt.Printf("Ultimo caractere: %c\n", ultimoCaracere)
}
