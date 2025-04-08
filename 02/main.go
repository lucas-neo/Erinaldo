package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"ListaDuplaRotas/rotas"
)

func main() {
	var continuar string = "s"

	for strings.ToLower(continuar) == "s" {
		fp, err := os.Open("rotas_de_transportes.txt")
		if err != nil {
			fmt.Println("Erro ao abrir o arquivo.")
			return
		}

		var filtro string
		fmt.Print("\nDigite um tipo de transporte ou regiao para buscar: ")
		fmt.Scanln(&filtro)

		var listaRotas *rotas.Rota = nil
		carregarRotasDeArquivo(&listaRotas, fp, filtro)
		fp.Close()

		if listaRotas == nil {
			fmt.Println("\nNenhuma rota encontrada para o filtro informado.")
		} else {
			for {
				fmt.Println("\n=== MENU DE ROTAS ===")
				fmt.Println("1 - Ver rotas em ordem normal")
				fmt.Println("2 - Ver rotas em ordem inversa")
				fmt.Println("3 - Fazer nova busca")
				fmt.Println("4 - Sair")
				fmt.Print("Escolha uma opção: ")

				var opcao int
				fmt.Scanln(&opcao)

				switch opcao {
				case 1:
					rotas.ExibirListaEmOrdem(listaRotas)
				case 2:
					rotas.ExibirListaEmOrdemInversa(listaRotas)
				case 3:
					rotas.LiberarLista(&listaRotas)
					listaRotas = nil
				case 4:
					rotas.LiberarLista(&listaRotas)
					listaRotas = nil
					continuar = "n"
				default:
					fmt.Println("Opção inválida.")
				}

				if listaRotas == nil || strings.ToLower(continuar) != "s" {
					break
				}
			}
		}

		if continuar == "s" {
			fmt.Print("\nDeseja fazer outra busca? (s/N): ")
			fmt.Scanln(&continuar)
		}
	}

	fmt.Println("\nPrograma encerrado.")
}

func carregarRotasDeArquivo(lista **rotas.Rota, fp *os.File, filtro string) {
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		linha := scanner.Text()
		partes := strings.SplitN(linha, ";", 4)
		if len(partes) != 4 {
			continue
		}

		id, err := strconv.Atoi(partes[0])
		if err != nil {
			continue
		}

		nome := partes[1]
		tipo := partes[2]
		regiao := partes[3]

		if strings.Contains(tipo, filtro) || strings.Contains(regiao, filtro) {
			rotas.InserirNaListaDupla(lista, id, nome, tipo, regiao)
		}
	}
}
