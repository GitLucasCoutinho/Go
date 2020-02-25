package main

import (
	"fmt"
)
/*
Em Go, esse conceito é implementado de forma particular:
- Campos privados (letra minúscula) ficam ocultos fora do pacote.
- Métodos públicos (letra maiúscula) são a interface de acesso controlado.
- O pacote funciona como a “fronteira” de visibilidade.
Portanto, quando você cria uma struct com campos não exportados e fornece funções/métodos exportados para manipular esses dados, está aplicando encapsulamento — ou seja, escondendo detalhes internos e expondo apenas o que é necessário.
Em outras linguagens orientadas a objetos, isso seria chamado de information hiding ou data hiding, mas em Go usamos o termo encapsulamento via exportação de identificadores
*/
// Go implementa encapsulamento através da visibilidade de identificadores e pacotes,
// sem precisar de modificadores de acesso tradicionais. Isso torna o modelo mais simples,
// mas igualmente eficaz para proteger dados e controlar o acesso.

// Definição de uma struct com campo privado
type ContaBancaria struct {
	saldo float64 // campo não exportado (privado ao pacote)
}

// Método público para criar uma nova conta
func NovaContaBancaria(saldoInicial float64) *ContaBancaria {
	return &ContaBancaria{saldo: saldoInicial}
}

// Método público para depositar dinheiro
func (c *ContaBancaria) Depositar(valor float64) {
	if valor > 0 {
		c.saldo += valor
	}
}

// Método público para sacar dinheiro
func (c *ContaBancaria) Sacar(valor float64) bool {
	if valor > 0 && valor <= c.saldo {
		c.saldo -= valor
		return true
	}
	return false
}

// Método público para consultar saldo
func (c *ContaBancaria) Saldo() float64 {
	return c.saldo
}

func main() {
	conta := NovaContaBancaria(100.0)

	conta.Depositar(50.0)
	fmt.Println("Saldo após depósito:", conta.Saldo())

	if conta.Sacar(30.0) {
		fmt.Println("Saque realizado com sucesso!")
	} else {
		fmt.Println("Saldo insuficiente.")
	}

	fmt.Println("Saldo final:", conta.Saldo())
}