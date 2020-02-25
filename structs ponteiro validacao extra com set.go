package main

import "fmt"

type Pessoa struct {
    nome  string
    idade int
}

// Método com RECEIVER por valor (não altera o original)
func (p Pessoa) SetIdadeValor(novaIdade int) {
    p.idade = novaIdade
}

// Método com RECEIVER por ponteiro (altera o original)
func (p *Pessoa) SetIdadePonteiro(novaIdade int) {
    p.idade = novaIdade
}

func main() {
    pessoa := Pessoa{nome: "Lucas", idade: 25}

    // Usando método por valor (cópia)
    pessoa.SetIdadeValor(30)
    fmt.Println("Após SetIdadeValor:", pessoa.idade) // continua 25

    // Usando método por ponteiro (referência)
    pessoa.SetIdadePonteiro(40)
    fmt.Println("Após SetIdadePonteiro:", pessoa.idade) // agora é 40
}