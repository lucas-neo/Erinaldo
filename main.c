#include <stdio.h>
#include <stdlib.h>

#define TAM 15
typedef struct no {
    int chave;
    struct no *proximo;
} No;

typedef struct {
    No *inicio;
    int tam;
}Lista;

void inicializarLista(Lista *l){
    l->inicio = NULL;
    l->tam = 0;
}


void inserir_lista(Lista *l, int valor){
    No *novo = malloc(sizeof(No));

    if(novo){
        novo->chave = valor;
        novo->proximo = l->inicio;
        l->inicio = novo;
        l->tam++;
    }else
    printf("Erro ao alocar memória.\n");
}

int buscar_lista(Lista *l, int valor){
    No *aux = l->inicio;
    while(aux && aux->chave != valor)
        aux=aux->proximo;
    if(aux)
        return aux->chave;
    return 0;
}

void imprimir_lista(Lista *l) {
    No *aux = l->inicio;
    printf("Tam: %d | ", l->tam);
    while (aux){
        printf("%d ", aux->chave);
        aux = aux->proximo;
    }
}

void inicializarTabela(Lista t[]) {
    int i;
    for (i = 0; i < TAM; i++)
        inicializarLista(&t[i]);
}

int funcaoHash(int chave){
    return chave % TAM;
}

void inserir(Lista t[], int valor) {
    int id = funcaoHash(valor);
    inserir_lista(&t[id], valor);    
}

int busca(Lista t[], int chave) {
    int id = funcaoHash(chave);
    return buscar_lista(&t[id], chave);
}

void imprimir(Lista t[]){
    int i;
    for(i = 0; i < TAM; i++){
        printf("%2d = ", i);
        imprimir_lista(&t[i]);
        printf("\n");
    }
}

int main () {
    int opcao, valor, retorno;
    Lista tabela[TAM];

    inicializarTabela(tabela);

    do{
        printf("\n0 - Sair\n1 - Inserir\n2 - Buscar\n3 - Imprimir\n");
        scanf("%d", &opcao);

        switch (opcao) {
            case 1:
                printf("\nValor: ");
                scanf("%d", &valor);
                inserir(tabela, valor);
                break;
            case 2:
                printf("\nValor pra buscar: ");
                scanf("%d", &valor);
                retorno = busca(tabela, valor);
                if (retorno != 0)
                    printf("\nValor encontrado: %d\n", retorno);
                else
                    printf("\nValor nao encontrado.\n");
                break;
            case 3:
                imprimir(tabela);
            break;
            default:
                printf("Opcao invalida");
        }
    }while(opcao != 0);

    return 0;
}
