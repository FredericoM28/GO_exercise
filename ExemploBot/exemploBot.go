package main

import (
	"fmt"
	// "math"
)

// Estrutura para os dados do mercado
type Tick struct {
	Simbolo string
	Preco   float64
}

// Analisador decide se entra na operação e calcula SL e TP
func analisarSinal(t Tick) (string, float64, float64) {
	// Aqui entraria sua lógica de IA ou Indicadores
	// Exemplo simples: Estratégia de reversão
	stopLoss := t.Preco - (t.Preco * 0.01)   // 1% abaixo
	takeProfit := t.Preco + (t.Preco * 0.02) // 2% acima

	return "COMPRA", stopLoss, takeProfit
}

func main() {
	// Simulando um fluxo de dados em tempo real (Canal)
	precos := make(chan Tick)

	// Goroutine que simula a corretora enviando dados
	go func() {
		tickers := []float64{1.0850, 1.0855, 1.0842, 1.0860} // Ex: EURUSD
		for _, p := range tickers {
			precos <- Tick{Simbolo: "EURUSD", Preco: p}
		}
		close(precos)
	}()

	fmt.Println(" Bot de Trading Go Iniciado...")
	fmt.Println("---------------------------------")

	for tick := range precos {
		sinal, sl, tp := analisarSinal(tick)
		fmt.Printf("[%s] Preço: %.4f | Sinal: %s | SL: %.4f | TP: %.4f\n", 
			tick.Simbolo, tick.Preco, sinal, sl, tp)
	}
}