package main

import (
    "fmt"
    "strings"
)

func main() {
    var palavra string

    fmt.Print("Digite uma palavra: ")
    fmt.Scan(&palavra)

    // Converte para minúsculas para facilitar a comparação
    palavra = strings.ToLower(palavra)

    vogais := "aeiou"
    contador := 0

    // Percorre cada caractere da palavra
    for _, letra := range palavra {
        if strings.ContainsRune(vogais, letra) {
            contador++
        }
    }

    fmt.Printf("A palavra '%s' possui %d vogais.\n", palavra, contador)
}