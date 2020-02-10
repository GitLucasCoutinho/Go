package main

import "fmt"

// Definição da struct Pessoa (equivalente a uma classe)
type Pessoa struct {
    nome  string
    idade int
}

// Método associado à struct Pessoa
func (p Pessoa) apresentar() {
    fmt.Printf("Olá, meu nome é %s e tenho %d anos.\n", p.nome, p.idade)
}

func main() {
    // Instanciando objetos do tipo Pessoa
    pessoa1 := Pessoa{nome: "Lucas", idade: 25}
    pessoa2 := Pessoa{nome: "Maria", idade: 30}

    // Chamando o método apresentar
    pessoa1.apresentar()
    pessoa2.apresentar()
}

/*
Saída esperada:

Olá, meu nome é Lucas e tenho 25 anos.
Olá, meu nome é Maria e tenho 30 anos.
*/