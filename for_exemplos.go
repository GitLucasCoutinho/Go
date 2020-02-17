package main

import "fmt"

func main() {
    // 1. For clássico: Tabuada do 7
    numero := 7
    fmt.Println("Tabuada do", numero)
    for i := 1; i <= 10; i++ {
        fmt.Printf("%d x %d = %d\n", numero, i, numero*i)
    }

    fmt.Println()

    // 2. For estilo while: contar até 5
    j := 1
    for j <= 5 {
        fmt.Println("Contando estilo while:", j)
        j++
    }

    fmt.Println()

    // 3. For infinito com break: imprime até 3
    k := 1
    for {
        fmt.Println("Loop infinito:", k)
        k++
        if k > 3 {
            break
        }
    }

    fmt.Println()

    // 4. For range: percorrendo string "Lucas"
    nome := "Lucas"
    fmt.Println("Percorrendo string:", nome)
    for i, letra := range nome {
        fmt.Printf("Posição %d -> %c\n", i, letra)
    }

    fmt.Println()

    // 5. For com continue: imprimir apenas ímpares até 10
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            continue
        }
        fmt.Println("Ímpar:", i)
    }

    fmt.Println()

    // 6. For com break: parar quando chegar em 5
    for i := 1; i <= 10; i++ {
        if i == 5 {
            break
        }
        fmt.Println("Até chegar no 5:", i)
    }
}