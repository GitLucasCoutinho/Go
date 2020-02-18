package main

import "fmt"

func main() {
    // Declarando um array de 10 números
    numeros := [10]int{5, 8, 12, 20, 7, 15, 9, 10, 18, 6}

    soma := 0
    for i := 0; i < len(numeros); i++ {
        soma += numeros[i]
    }

    media := float64(soma) / float64(len(numeros))

    fmt.Println("Array:", numeros)
    fmt.Println("Soma:", soma)
    fmt.Println("Média:", media)
}