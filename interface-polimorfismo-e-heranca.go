package main

import "fmt"

// Interface define um comportamento comum
type Animal interface {
    Falar() string
}

// Struct base (não há herança, mas podemos compor da seguinte forma)
type SerVivo struct {
    nome string
}

// Structs específicas que "herdam" SerVivo por composição
type Cachorro struct {
    SerVivo
}

type Gato struct {
    SerVivo
}

// Implementação dos métodos da interface
func (c Cachorro) Falar() string {
    return "Au au! Eu sou " + c.nome
}

func (g Gato) Falar() string {
    return "Miau! Eu sou " + g.nome
}

// Função que demonstra polimorfismo
func FazerAnimalFalar(a Animal) {
    fmt.Println(a.Falar())
}

func main() {
    // Criando instâncias
    cachorro := Cachorro{SerVivo{"Rex"}}
    gato := Gato{SerVivo{"Mimi"}}

    // Polimorfismo: ambos são "Animal"
    FazerAnimalFalar(cachorro)
    FazerAnimalFalar(gato)
}