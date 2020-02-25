package main

import "fmt"

type Pessoa struct {
    nome  string
    idade int
}

// Setter com validação
func (p *Pessoa) SetIdade(novaIdade int) {
    if novaIdade >= 0 {
        p.idade = novaIdade
    } else {
        fmt.Println("Idade inválida!")
    }
}

// Getter
func (p Pessoa) GetIdade() int {
    return p.idade
}

func main() {
    pessoa := Pessoa{nome: "Lucas"}
    pessoa.SetIdade(25) // válido
    fmt.Println("Idade:", pessoa.GetIdade())

    pessoa.SetIdade(-5) // inválido
}