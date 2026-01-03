package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

// Bloco com suporte a Mineração (Nonce)
type Bloco struct {
	Indice       int
	Timestamp    string
	Dados        string
	HashAnterior string
	HashAtual    string
	Nonce        int // O "número mágico" da mineração
}

// CalcularHash combina os dados e o Nonce para gerar a digital do bloco
func calcularHash(b Bloco) string {
	registro := fmt.Sprintf("%d%s%s%s%d", b.Indice, b.Timestamp, b.Dados, b.HashAnterior, b.Nonce)
	h := sha256.New()
	h.Write([]byte(registro))
	return hex.EncodeToString(h.Sum(nil))
}

// minerarBloco faz o computador trabalhar até achar um hash que comece com N zeros
func minerarBloco(b *Bloco, dificuldade int) {
	alvo := strings.Repeat("0", dificuldade)
	for {
		b.HashAtual = calcularHash(*b)
		if strings.HasPrefix(b.HashAtual, alvo) {
			fmt.Printf("⛏️  Bloco Minerado! Hash: %s\n", b.HashAtual)
			break
		}
		b.Nonce++ // Tenta o próximo número
	}
}

func main() {
	dificuldade := 4 // Tente aumentar para 5 ou 6 e veja o PC esquentar!

	// 1. Criando o Bloco Gênesis usando Campos Nomeados (Evita o erro que você teve)
	genesis := Bloco{
		Indice:       0,
		Timestamp:    time.Now().Format("02-01-2006 15:04:05"),
		Dados:        "Bloco Gênesis (Origem)",
		HashAnterior: "0",
		Nonce:        0,
	}
	
	fmt.Println("Minerando bloco gênesis...")
	minerarBloco(&genesis, dificuldade)

	blockchain := []Bloco{genesis}

	// 2. Adicionando um novo bloco de transação de mercado
	novoBloco := Bloco{
		Indice:       1,
		Timestamp:    time.Now().Format("02-01-2006 15:04:05"),
		Dados:        "Sinal IA: COMPRA XAUUSD @ 2035.50",
		HashAnterior: genesis.HashAtual,
		Nonce:        0,
	}

	fmt.Println("Minerando bloco de sinal de trading...")
	minerarBloco(&novoBloco, dificuldade)
	blockchain = append(blockchain, novoBloco)

	fmt.Println("\n--- CADEIA DE BLOCOS FINAL ---")
	for _, b := range blockchain {
		fmt.Printf("Bloco #%d | Hash: %s...\n", b.Indice, b.HashAtual[:15])
	}
}