package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"ListaCircular/rodizio"
)

const Arquivo = "rodizio_de_veiculos.txt"

func main() {
	var continuar string = "s"
	leitor := bufio.NewReader(os.Stdin)

	for strings.ToLower(continuar) == "s" {
		var lista rodizio.Lista
		rodizio.CriarLista(&lista)

		fp, err := os.Open(Arquivo)
		if err != nil {
			fmt.Printf("Erro ao abrir o arquivo '%s'.\n", Arquivo)
			return
		}

		fmt.Print("\nInforme o dia da semana (ex: Segunda-feira): ")
		filtro, _ := leitor.ReadString('\n')
		filtro = strings.TrimSpace(filtro)

		rodizio.CarregarVeiculos(&lista, fp, filtro)
		fp.Close()

		if lista.Inicio == nil {
			fmt.Println("Nenhum veículo encontrado para o filtro informado.")
		} else {
			escolha := "s"
			for strings.ToLower(escolha) == "s" {
				fmt.Print("\nDeseja:\n1 - Exibir todos os veículos\n2 - Navegar veículo por veículo\nEscolha: ")

				var opcao int
				fmt.Scanln(&opcao)

				switch opcao {
				case 1:
					rodizio.ExibirListaCircular(&lista)
				case 2:
					rodizio.NavegarPelaLista(&lista)
				default:
					fmt.Println("Opção inválida.")
				}

				fmt.Print("\nDeseja ver novamente ou navegar de novo com o mesmo filtro? (s/N): ")
				escolha, _ = leitor.ReadString('\n')
				escolha = strings.TrimSpace(escolha)
			}
		}

		rodizio.LiberarListaCircular(&lista)

		fmt.Print("\nDeseja realizar nova busca com outro filtro? (s/N): ")
		continuar, _ = leitor.ReadString('\n')
		continuar = strings.TrimSpace(continuar)
	}

	fmt.Println("\nPrograma encerrado. Até logo!")
}
