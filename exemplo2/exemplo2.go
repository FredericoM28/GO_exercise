package main
import "fmt"
func main(){
	fmt.Println("Digite seu nome: ")
	var nome string //declaracao da variavel nome

	fmt.Scanf("%s", &nome)// armazena o nome do usuario na var "nome"
	var serie string // outra forma " := "
	serie = "ao Curso de programacao com linguagem Go"
	//serie := "ao curso de programacao com linguagem Go" // outra forma para declarar variaveis e' usando " := "
	fmt.Printf("seja bem-vindo(a), %s, %s!",serie, nome)
}