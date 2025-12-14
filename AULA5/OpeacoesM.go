package main

import "fmt"

func main() {
	fmt.Println("Digite o Primeiro valor:")
	var num1 int
	fmt.Scanf("%d\n",&num1)

	fmt.Println("Digite o Segundo valor:")
	var num2 int
	fmt.Scanf("%d\n",&num2)

	soma:= num1 + num2
	subtracao := num1 - num2
	multiplicacao := num1 * num2
	divisao := num1 / num2
	modulo:= num1 % num2

	fmt.Printf("%d + %d = %d\n", num1, num1, soma)
	fmt.Printf("%d - %d = %d\n", num1, num1, subtracao)
	fmt.Printf("%d * %d = %d\n", num1, num1, multiplicacao)
	fmt.Printf("%d / %d = %d\n", num1, num1, divisao)
	fmt.Printf("%d %% %d = %d\n", num1, num1, modulo)
}