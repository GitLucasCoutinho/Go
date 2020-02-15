package main

import "fmt"

func main() {
    var numero int

    fmt.Print("Digite um número: ")
    fmt.Scan(&numero) // lê o valor digitado pelo usuário

    if numero%2 == 0 {
        fmt.Println("O número é par.")
    } else {
        fmt.Println("O número é ímpar.")
    }
}