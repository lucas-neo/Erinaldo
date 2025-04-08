package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"listaPontosColeta/coleta"
)

func main() {
	fp, err := os.Open("pontos_de_descartes.txt") //como funciona esse os.Open, por que estou colocando duas variaveis?
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo.")
		return
	}
	defer fp.Close() //explica melhor sobre o defer

	processarBuscas(fp)

	fmt.Println("\nSaindo...")
}

func processarBuscas(fp *os.File) {
	leitor := bufio.NewReader(os.Stdin) //da onde veio o os.Stdin?
	var continuar string

	for {
		fmt.Print("Qual material que deseja descartar: ")
		materialEntrada, erroMaterial := leitor.ReadString('\n') //mesma coisa aqui, por que estou colocando duas variaveis?
		if erroMaterial != nil {
			fmt.Println("Erro na leitura do material.")
			return
		}
		material := strings.TrimSpace(materialEntrada) //por usar o TrimSpace?

		var lista *coleta.PontoColeta = nil

		fp.Seek(0, 0)
		scanner := bufio.NewScanner(fp) //o que faz?

		for scanner.Scan() {
			linha := scanner.Text()
			partes := strings.SplitN(linha, ";", 3)
			if len(partes) != 3 {
				fmt.Printf("Erro ao ler a linha: %s\n", linha)
				continue
			}

			id, erroConversao := strconv.Atoi(partes[0])
			if erroConversao != nil {
				continue
			}

			tipo := partes[1]
			endereco := partes[2]

			if tipo == material {
				coleta.InserirPontoColeta(&lista, id, tipo, endereco)
			}
		}

		if lista != nil {
			coleta.ListarPontosColeta(lista)
		} else {
			fmt.Printf("Nenhum ponto de coleta encontrado para o material %s.\n", material)
		}

		coleta.LiberarLista(&lista)

		fmt.Print("\nContinuar? (y/n): ")
		resposta, erroLeitura := leitor.ReadString('\n')
		if erroLeitura != nil {
			fmt.Println("Erro na leitura da resposta.")
			return
		}
		continuar = strings.TrimSpace(resposta)
		if strings.ToLower(continuar) != "y" {
			break
		}
	}
}
