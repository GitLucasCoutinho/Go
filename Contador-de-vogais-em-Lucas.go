package main // Pacote principal

import "fmt" // Importa o pacote fmt para imprimir no console

func main() { // Função principal
    nome := "Lucas" // Variável string com valor "Lucas"

    // Laço for com range: percorre cada caractere da string
    // 'i' recebe o índice (posição do caractere)
    // 'letra' recebe o caractere em formato rune
    for i, letra := range nome {
        // %d -> imprime número decimal (índice)
        // %c -> imprime caractere
        fmt.Printf("Índice: %d -> Letra: %c\n", i, letra)
    }
}

/*
Saída esperada:

Índice: 0 -> Letra: L
Índice: 1 -> Letra: u
Índice: 2 -> Letra: c
Índice: 3 -> Letra: a
Índice: 4 -> Letra: s
*/