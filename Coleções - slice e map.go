package main

import "fmt"

func main() {
    // Slice: coleção ordenada (similar a lista)
    frutas := []string{"maçã", "banana", "laranja"}
    fmt.Println("Coleção de frutas (slice):", frutas)

    // Adicionando elementos ao slice
    frutas = append(frutas, "uva")
    fmt.Println("Após adicionar uva:", frutas)

    // Iterando sobre o slice (Um slice é uma estrutura que representa uma sequência dinâmica de elementos)
)
    fmt.Println("Iterando sobre frutas:")
    for i, fruta := range frutas {
        fmt.Printf("Índice %d: %s\n", i, fruta)
    }

    // Map: coleção chave-valor (similar a dicionário)
    capitais := map[string]string{
        "Brasil":   "Brasília",
        "França":   "Paris",
        "Japão":    "Tóquio",
    }
    fmt.Println("\nColeção de capitais (map):", capitais)

    // Adicionando elemento ao map
    capitais["Alemanha"] = "Berlim"

    // Iterando sobre o map
    fmt.Println("Iterando sobre capitais:")
    for pais, capital := range capitais {
        fmt.Printf("País: %s -> Capital: %s\n", pais, capital)
    }

    // Verificando se uma chave existe
    if capital, existe := capitais["Brasil"]; existe {
        fmt.Println("\nCapital do Brasil é:", capital)
    } else {
        fmt.Println("\nBrasil não encontrado no mapa.")
    }
}