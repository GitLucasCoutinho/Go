package main

import "fmt"

// Função que recebe dois números e retorna o maior
func maior(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    x := 10
    y := 25

    resultado := maior(x, y)
    fmt.Printf("Entre %d e %d, o maior é %d\n", x, y, resultado)
}