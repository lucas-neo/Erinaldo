package coleta

import "fmt"

type PontoColeta struct {
	ID           int
	TipoMaterial string
	Endereco     string
	Prox         *PontoColeta
}

func InserirPontoColeta(lista **PontoColeta, id int, tipoMaterial string, endereco string) {
	novo := &PontoColeta{
		ID:           id,
		TipoMaterial: tipoMaterial,
		Endereco:     endereco,
		Prox:         nil,
	}

	if *lista == nil {
		*lista = novo
	} else {
		atual := *lista
		for atual.Prox != nil {
			atual = atual.Prox
		}
		atual.Prox = novo
	}
}

func ListarPontosColeta(lista *PontoColeta) {
	if lista == nil {
		fmt.Println("Nenhum ponto de coleta encontrado.")
		return
	}

	fmt.Printf("PONTOS DE COLETA PARA: %s\n", lista.TipoMaterial)
	fmt.Println("----------------------------------------------------")
	atual := lista
	for atual != nil {
		fmt.Printf("ID: %-3d | Tipo: %-6s | Endereço: %-30s\n", atual.ID, atual.TipoMaterial, atual.Endereco)
		fmt.Println("----------------------------------------------------")
		atual = atual.Prox
	}
}

func LiberarLista(lista **PontoColeta) {
	*lista = nil // Garbage Collector cuida da liberação
}
