package main

import "fmt"

func main() {
    // Declarando duas variáveis
    var num1 float64 = 10
    var num2 float64 = 5

    // Operações matemáticas
    soma := num1 + num2
    subtracao := num1 - num2
    multiplicacao := num1 * num2
    divisao := num1 / num2

    // Imprimindo os resultados
    fmt.Println("Número 1:", num1)
    fmt.Println("Número 2:", num2)
    fmt.Println("Soma:", soma)
    fmt.Println("Subtração:", subtracao)
    fmt.Println("Multiplicação:", multiplicacao)
    fmt.Println("Divisão:", divisao)
}