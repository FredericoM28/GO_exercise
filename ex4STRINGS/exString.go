package main
import (
	"fmt"
	"Strings"
)

func main (){
	frase := "Um delicioso bolo de laranja"
	fmt.Printf("Frse original: %s\n", frase)
	fraseMinuscula := strings.ToLower(frase)
	fmt.Printf("Frase minuscula: %s\n", fraseMinuscula)	
	fraseMaiuscula := strings.ToUpper(frase)
	fmt.Printf("Frase maaiuscula: %s\n", fraseMaiuscula)

	outraFrase:= "    uma frase no meio     "
	tamanho := len(outraFrase)
	fmt.Printf("tamanho da frase: %d\n", tamanho)
	outraFrase = strings.Trim(outraFrase, " ")
	tamanho = len(outraFrase)
	fmt.Printf("tamanho da frase depois: %d\n", tamanho)


}