package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Resultado define a estrutura da nossa resposta JSON
type Resultado struct {
	URL    string `json:"url"`
	Status string `json:"status"`
	Tempo  string `json:"tempo_resposta"`
}

// verificarSite tenta acessar uma URL e envia o resultado para um canal (channel)
func verificarSite(url string, c chan Resultado) {
	inicio := time.Now()
	resp, err := http.Get(url)
	duracao := time.Since(inicio).String()

	if err != nil || resp.StatusCode != 200 {
		c <- Resultado{URL: url, Status: "FORA DO AR ", Tempo: duracao}
		return
	}
	c <- Resultado{URL: url, Status: "ONLINE ", Tempo: duracao}
}

func monitorHandler(w http.ResponseWriter, r *http.Request) {
	sites := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.google.com.br",
	}

	canal := make(chan Resultado)
	var resultados []Resultado

	// Dispara a verificação de cada site em paralelo usando Goroutines
	for _, url := range sites {
		go verificarSite(url, canal)
	}

	// Coleta os resultados assim que eles chegam no canal
	for i := 0; i < len(sites); i++ {
		resultados = append(resultados, <-canal)
	}

	// Retorna os dados como JSON para o navegador/cliente
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultados)
}

func main() {
	http.HandleFunc("/status", monitorHandler)

	fmt.Println("Servidor rodando em http://localhost:8080/status")
	http.ListenAndServe(":8080", nil)
}