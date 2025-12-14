package main

import "fmt"

func main(){
	resultado1 := 5
	resultado2 := 7
	
	
	resultado1++
	resultado2--
	fmt.Printf("Resutado 1: %d\n", resultado1)
	fmt.Printf("Resutado 2: %d\n", resultado2)

	resultado1 = 5 + 2 *10
	resultado2 = (5 + 2) * 10
	fmt.Printf("Resutado 1: %d\n", resultado1)
	fmt.Printf("Resutado 2: %d\n", resultado2)

}
