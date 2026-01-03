package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// --- TAREFA 1: ORGANIZADOR DE ARQUIVOS ---
func organizadorTask() {
	pasta := "./Downloads" // Altere para sua pasta de downloads
	for {
		files, _ := os.ReadDir(pasta)
		for _, f := range files {
			if filepath.Ext(f.Name()) == ".exe" || filepath.Ext(f.Name()) == ".msi" {
				fmt.Printf("🧹 Sentinel: Encontrei instalador %s, movendo...\n", f.Name())
				// Lógica de mover arquivo entraria aqui
			}
		}
		time.Sleep(1 * time.Hour) // Roda a cada hora
	}
}

// --- TAREFA 2: MONITOR DE MERCADO (SIMULADO) ---
func monitorMercadoTask() {
	for {
		// Aqui você faria um http.Get em uma API de Forex/Gold
		precoOuro := 2045.50 // Simulação
		if precoOuro > 2040 {
			fmt.Printf("📈 ALERTA: XAUUSD está em alta: $%.2f\n", precoOuro)
		}
		time.Sleep(30 * time.Second)
	}
}

// --- TAREFA 3: INTERFACE DE CONTROLE (API) ---
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>🛡️ Go-Sentinel Ativo</h1><p>Monitorando Mercado e Arquivos...</p>")
}

func main() {
	// Inicializa as tarefas em background (Goroutines)
	go organizadorTask()
	go monitorMercadoTask()

	// Servidor para você ver o status no navegador
	http.HandleFunc("/", homeHandler)

	fmt.Println(`
    	  ____           ____             _   _             _ 
    	 / ___| ___     / ___|  ___ _ __ | |_(_)_ __   ___| |
    	| |  _ / _ \    \___ \ / _ \ '_ \| __| | '_ \ / _ \ |
    	| |_| | (_) |    ___) |  __/ | | | |_| | | | |  __/ |
    	 \____|\___/    |____/ \___|_| |_|\__|_|_| |_|\___|_|
	`)
	fmt.Println("🚀 Sentinel rodando em http://localhost:9000")
	
	log.Fatal(http.ListenAndServe(":9000", nil))
}