package main

import "fmt"

func main() {
    // Declarando variáveis de diferentes tipos
    var nome string = "Lucas"     // string
    var idade int = 25            // inteiro
    var altura float64 = 1.78     // número decimal (float64)
    var ativo bool = true         // booleano
    var inicial rune = 'L'        // caractere (rune)
    var nota byte = 65            // byte (uint8)

    // Imprimindo os valores
    fmt.Println("Nome:", nome)
    fmt.Println("Idade:", idade)
    fmt.Println("Altura:", altura)
    fmt.Println("Ativo:", ativo)
    fmt.Println("Inicial:", inicial)
    fmt.Println("Nota (byte):", nota)
}