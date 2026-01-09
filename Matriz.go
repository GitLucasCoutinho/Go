package main

import "fmt"

func main() {
    // Declarando uma matriz 3x3
    matriz := [3][3]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }

    fmt.Println("Matriz 3x3:")
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            fmt.Printf("%d ", matriz[i][j])
        }
        fmt.Println()
    }
}