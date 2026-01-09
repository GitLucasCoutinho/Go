
Este repositório reúne diversos exemplos práticos em **Go (Golang)**, explorando conceitos fundamentais da linguagem como variáveis, estruturas de controle, funções, structs, interfaces, polimorfismo, encapsulamento, goroutines (threads), tratamento de erros e coleções.  
O objetivo é servir como um guia de estudo e prática para iniciantes e intermediários em Go.

---

## 📂 Estrutura do Projeto

### 🔹 Conceitos básicos
- **Olá-mundo.go** → Primeiro programa em Go.
- **Variaveis e tipos.go** → Declaração de variáveis e tipos básicos.
- **Operações-Matemáticas.go** → Operações aritméticas simples.
- **Condicional-simples.go** → Estruturas condicionais (`if`, `else`).
- **for_exemplos.go** → Exemplos e variações do laço `for`.

### 🔹 Funções e modularização
- **funcoes.go** → Definição e uso de funções.
- **Contador-de-vogais.go** → Função que conta vogais em uma string.
- **Contador-de-vogais-em-Lucas.go** → Exemplo aplicado ao nome "Lucas".

### 🔹 Arrays, slices e maps (coleções)
- **Array.go** → Uso de arrays fixos.
- **Matriz.go** → Manipulação de matrizes.
- **Coleções - slice e map.go** → Uso de slices (listas dinâmicas) e maps (dicionários).

### 🔹 Structs e encapsulamento
- **structs.go** → Definição básica de structs.
- **structs-e-construtor.go** → Construtores em Go.
- **structs exemplo simples com ponteiro.go** → Uso de ponteiros em structs.
- **structs acesso direto com ponteiro.go** → Acesso direto com ponteiros.
- **structs ponteiro validacao extra com set.go** → Validação com métodos `set`.
- **encapsulamento.go** → Encapsulamento via visibilidade de identificadores.

### 🔹 Interfaces, polimorfismo e herança (composição)
- **interface-polimorfismo-e-heranca.go** → Demonstração de interfaces e polimorfismo.
- **Formas de comentarios.go** → Exemplos de comentários e documentação.
- **interface polimorfismo e heranca** → Outro exemplo prático de composição e polimorfismo.

### 🔹 Concorrência
- **goroutines a thread do Go.go** → Introdução às goroutines.
- **goroutines (threads).go** → Exemplo de execução concorrente (números e letras).

### 🔹 Tratamento de erros
- **tratar erros com errors dot New.go** → Uso de `errors.New` para retornar erros.
- **tratar erros com panic e recover.go** → Simulação de try-catch com `panic` e `recover`.

---

## 🚀 Como executar

1. Instale o Go: [https://go.dev/dl/](https://go.dev/dl/)
2. Clone este repositório:
   ```bash
   git clone https://github.com/GitLucasCoutinho/Go-Exemplos.git

   cd Go-Exemplos
   go run Olá-mundo.go
