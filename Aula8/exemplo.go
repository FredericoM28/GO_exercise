package main

import "fmt"

func main() {
	fmt.Printf("Digite um numero inteiro positivo: ")
	var n int
	fmt.Scanf("%d", &n)

	if n < 0 {
		n *= -1
	} else if n == 0{
		n = 1
	}

	// LOOP normal
	/*contaddor := 1
	for contaddor <= n{
		fmt.Printf("%d", contaddor)
		contaddor = contaddor + 1
	}*/

   // LOOP reduzido	normal de 1 ate n
	 for contaddor := 1; contaddor <= n;  contaddor++ {
		fmt.Printf("%d", contaddor)
		
	} 

	//loop inverso de (n ate 1)
	for contador := n; contador >=1; contador--{
		fmt.Printf("%d", contador)

	}

}
