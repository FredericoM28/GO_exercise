package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// 1. Cria a aplicação e a janela principal
	meuApp := app.New()
	janela := meuApp.NewWindow("GoTrade - IA Bot")
	janela.Resize(fyne.NewSize(400, 300))

	// 2. Cria elementos da interface (Widgets)
	titulo := widget.NewLabel("Monitor de Mercado - XAUUSD")
	precoLabel := widget.NewLabel("Preço Atual: $2035.50")
	precoLabel.TextStyle = fyne.TextStyle{Bold: true}

	resultado := widget.NewLabel("Aguardando análise da IA...")

	// 3. Botão para acionar a lógica do Bot
	botaoAnalise := widget.NewButton("Gerar Sinal IA", func() {
		// Aqui você chamaria sua função de IA
		resultado.SetText("SINAL: COMPRA\nSL: 2030.00 | TP: 2050.00")
		fmt.Println("Botão clicado: Analisando mercado...")
	})

	// 4. Organiza tudo em um layout (Vertical)
	conteudo := container.NewVBox(
		titulo,
		precoLabel,
		widget.NewSeparator(),
		resultado,
		botaoAnalise,
	)

	// 5. Exibe a janela e roda o app
	janela.SetContent(conteudo)
	janela.ShowAndRun()
}