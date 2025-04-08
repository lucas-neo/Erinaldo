package rodizio

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type No struct {
	PlacaVeiculo string
	DiaDaSemana  string
	Horario      string
	Proximo      *No
}

type Lista struct {
	Inicio *No
	Fim    *No
}

func CriarLista(lista *Lista) {
	lista.Inicio = nil
	lista.Fim = nil
}

func InserirNaListaCircular(lista *Lista, placa, dia, horario string) {
	novo := &No{
		PlacaVeiculo: placa,
		DiaDaSemana:  dia,
		Horario:      horario,
		Proximo:      nil,
	}

	if lista.Inicio == nil {
		lista.Inicio = novo
		lista.Fim = novo
		novo.Proximo = novo // circular
	} else {
		novo.Proximo = lista.Inicio
		lista.Fim.Proximo = novo
		lista.Fim = novo
	}
}

func CarregarVeiculos(lista *Lista, fp *os.File, filtro string) {
	scanner := bufio.NewScanner(fp)
	contador := 0

	for scanner.Scan() {
		linha := scanner.Text()
		partes := strings.SplitN(linha, ";", 3)
		if len(partes) != 3 {
			fmt.Printf("Erro ao ler a linha: %s\n", linha)
			continue
		}

		placa := strings.TrimSpace(partes[0])
		dia := strings.TrimSpace(partes[1])
		horario := strings.TrimSpace(partes[2])

		if dia == filtro {
			InserirNaListaCircular(lista, placa, dia, horario)
			contador++
		}
	}

	if contador == 0 {
		fmt.Println("\nNenhum veículo encontrado para o filtro informado.")
	}
}

func ExibirListaCircular(lista *Lista) {
	if lista.Inicio == nil {
		fmt.Println("\nLista vazia.")
		return
	}

	fmt.Println("\nVEÍCULOS AFETADOS PELO RODÍZIO:")
	fmt.Println("Placa       | Dia da Semana     | Horário")
	fmt.Println("---------------------------------------------")

	atual := lista.Inicio
	for {
		fmt.Printf("%-12s| %-18s| %s\n", atual.PlacaVeiculo, atual.DiaDaSemana, atual.Horario)
		atual = atual.Proximo
		if atual == lista.Inicio {
			break
		}
	}
}

func NavegarPelaLista(lista *Lista) {
	if lista.Inicio == nil {
		fmt.Println("\nLista vazia.")
		return
	}

	atual := lista.Inicio
	leitor := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("\nPLACA: %s\nDIA: %s\nHORÁRIO: %s\n", atual.PlacaVeiculo, atual.DiaDaSemana, atual.Horario)
		fmt.Print("\nDeseja ver o próximo veículo? (s/N): ")
		opcao, _ := leitor.ReadString('\n')
		opcao = strings.TrimSpace(opcao)

		if strings.ToLower(opcao) != "s" {
			break
		}

		atual = atual.Proximo
	}
}

func LiberarListaCircular(lista *Lista) {
	lista.Inicio = nil
	lista.Fim = nil
}
