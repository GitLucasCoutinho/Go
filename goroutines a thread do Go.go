package main

import (
	"fmt"
	"sync"
)

/*
Este programa demonstra o uso de goroutines em Go.
- Uma goroutine imprime números de 1 a 50.
- A função principal imprime letras de A a Z.
- Usamos sync.WaitGroup para sincronizar a execução.
*/

// Função que imprime números de 1 a 50
func imprimirNumeros(wg *sync.WaitGroup) {
	defer wg.Done() // garante que o WaitGroup seja decrementado ao final
	for i := 1; i <= 50; i++ {
		fmt.Println("Número:", i) // comentário de linha única
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1) // adiciona 1 tarefa ao WaitGroup

	// Inicia a goroutine (thread leve)
	go imprimirNumeros(&wg)

	// Enquanto isso, a função principal imprime letras
	for ch := 'A'; ch <= 'Z'; ch++ {
		fmt.Println("Letra:", string(ch)) // imprime cada letra
	}

	// Espera a goroutine terminar antes de encerrar o programa
	wg.Wait()
}