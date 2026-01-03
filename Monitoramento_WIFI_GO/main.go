package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"sync"
)

// Função que executa ping em um IP
func ping(ip string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("ping", "-n", "1", "-w", "500", ip)
	} else {
		cmd = exec.Command("ping", "-c", "1", "-W", "1", ip)
	}

	err := cmd.Run()
	if err == nil {
		results <- ip
	}
}

func main() {
	network := "192.168.1" // 🔴 ajuste conforme sua rede
	var wg sync.WaitGroup
	results := make(chan string, 254)

	fmt.Println("🔍 Escaneando dispositivos na rede...")

	for i := 1; i <= 254; i++ {
		ip := fmt.Sprintf("%s.%d", network, i)
		wg.Add(1)
		go ping(ip, &wg, results)
	}

	// Fecha o canal quando tudo terminar
	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("\n📡 Dispositivos encontrados:")
	for ip := range results {
		fmt.Println("✅", ip)
	}
}
