package main

import "fmt"

func dividir(a, b int) int {
	if b == 0 {
		panic("divisão por zero!") // dispara exceção
	}
	return a / b
}

func main() {
	// Simula try-catch com defer + recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Erro capturado:", r)
		}
	}()

	fmt.Println("Resultado:", dividir(10, 0))
}