package main

import (
	"errors"
	"fmt"
)

// Função que trata divisão e retorna erro se divisor for zero
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero não permitida")
	}
	return a / b, nil
}

func main() {
	// Caso válido
	resultado, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}

	// Caso com divisão por zero
	resultado, err = dividir(10, 0)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}
}