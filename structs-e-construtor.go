package main

import "fmt"

// Definição da struct Pessoa
type Pessoa struct {
    nome  string
    idade int
}

// Construtor: função que cria e retorna uma nova Pessoa
func NovaPessoa(nome string, idade int) Pessoa {
    return Pessoa{nome: nome, idade: idade}
}

// Método associado à struct Pessoa
func (p Pessoa) apresentar() {
    fmt.Printf("Olá, meu nome é %s e tenho %d anos.\n", p.nome, p.idade)
}

func main() {
    // Usando o construtor para instanciar objetos
    pessoa1 := NovaPessoa("Lucas", 25)
    pessoa2 := NovaPessoa("Maria", 30)

    pessoa1.apresentar()
    pessoa2.apresentar()
}

/*
Saída esperada:

Olá, meu nome é Lucas e tenho 25 anos.
Olá, meu nome é Maria e tenho 30 anos.
*/