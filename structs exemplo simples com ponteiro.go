package main

import "fmt"

type Pessoa struct {
    nome  string
    idade int
}

// Setter usando ponteiro (*Pessoa)
func (p *Pessoa) SetIdade(novaIdade int) {
    p.idade = novaIdade
}

func main() {
    pessoa := Pessoa{nome: "Lucas", idade: 25}

    fmt.Println("Antes:", pessoa.idade)

    pessoa.SetIdade(30) // altera direto o objeto original

    fmt.Println("Depois:", pessoa.idade)
}