// condicoes Logicas
package main
 
import "fmt"

func main () {
	fmt.Printf("Digite o dividendo: ")
	var dividendo float64
	fmt.Scanf("%f", &dividendo)

	fmt.Printf("Digite o divisor: ")
	var divisor float64
	fmt.Scanf("%f", &divisor)

	if divisor != 0 {
		resultado :=  dividendo / divisor
		fmt.Printf("%f / %f = %f\n", dividendo, divisor, resultado)
	} else {
		fmt.Printf("Divisor nao pode ser igual a 0")
	}
}