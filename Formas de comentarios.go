package main

import "fmt"

// Este é um comentário de linha única.
// Ele serve para explicar rapidamente algo no código.
// Você pode usar vários "//" seguidos para criar blocos de explicação.

/*
Este é um comentário de múltiplas linhas.
Ele é útil quando você precisa escrever explicações maiores,
documentar trechos de código ou desativar blocos temporariamente.
Você pode escrever quantas linhas quiser até fechar com "*/".
*/

func main() {
    fmt.Println("Demonstração de comentários em Go")

    // Comentário de linha única dentro da função main
    fmt.Println("Linha única: usando //")

    /*
       Comentário de múltiplas linhas dentro da função main
       Aqui podemos escrever várias observações
       sem precisar repetir // em cada linha.
    */
    fmt.Println("Múltiplas linhas: usando /* ... */")
}