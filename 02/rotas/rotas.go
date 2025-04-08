package rotas

import "fmt"

type Rota struct {
	ID             int
	NomeDaLinha    string
	TipoTransporte string
	Regiao         string
	Anterior       *Rota
	Proximo        *Rota
}

func CriarNo(id int, nomeDaLinha, tipoTransporte, regiao string) *Rota {
	return &Rota{
		ID:             id,
		NomeDaLinha:    nomeDaLinha,
		TipoTransporte: tipoTransporte,
		Regiao:         regiao,
		Anterior:       nil,
		Proximo:        nil,
	}
}

func InserirNaListaDupla(lista **Rota, id int, nomeDaLinha, tipoTransporte string, regiao string) {
	novo := CriarNo(id, nomeDaLinha, tipoTransporte, regiao)

	if *lista == nil {
		*lista = novo
	} else {
		atual := *lista
		for atual.Proximo != nil {
			atual = atual.Proximo
		}
		atual.Proximo = novo
		novo.Anterior = atual
	}
}

func ExibirListaEmOrdem(lista *Rota) {
	if lista == nil {
		fmt.Println("Nenhuma rota encontrada.")
		return
	}

	fmt.Println("\nROTA - ORDEM NORMAL:")
	fmt.Println("ID   | Linha                | Tipo        | Região")
	fmt.Println("-----------------------------------------------------")

	atual := lista
	for atual != nil {
		fmt.Printf("%-4d | %-20s | %-10s | %s\n", atual.ID, atual.NomeDaLinha, atual.TipoTransporte, atual.Regiao)
		atual = atual.Proximo
	}
}

func ExibirListaEmOrdemInversa(lista *Rota) {
	if lista == nil {
		fmt.Println("Nenhuma rota encontrada.")
		return
	}

	atual := lista
	for atual.Proximo != nil {
		atual = atual.Proximo
	}

	fmt.Println("\nROTA - ORDEM INVERSA:")
	fmt.Println("ID   | Linha                | Tipo        | Região")
	fmt.Println("-----------------------------------------------------")

	for atual != nil {
		fmt.Printf("%-4d | %-20s | %-10s | %s\n", atual.ID, atual.NomeDaLinha, atual.TipoTransporte, atual.Regiao)
		atual = atual.Anterior
	}
}

func LiberarLista(lista **Rota) {
	*lista = nil // GC cuida da memória
}
