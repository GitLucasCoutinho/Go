# 🐹 Roteiro de Estudo de Go (Do Básico ao Avançado)


## 🔹 Fundamentos Essenciais
- **Instalação e ambiente**
  - Instale o Go via [site oficial](https://go.dev).
  - Configure `GOPATH` e `GOROOT`.
  - Use `go run`, `go build`, `go fmt`, `go mod init`.
- **Sintaxe básica**
  - Tipos primitivos: `int`, `float64`, `string`, `bool`.
  - Declaração de variáveis (`var`, `:=`).
  - Constantes (`const`).
- **Controle de fluxo**
  - `if/else`, `switch`, `for` (único loop).
  - `break`, `continue`.
- **Funções**
  - Parâmetros, retorno múltiplo.
  - Funções anônimas e closures.
- **Exercícios**
  - Programa que soma dois números.
  - Verificador de número par/ímpar.
  - Conversor de temperatura (Celsius ↔ Fahrenheit).

---

## 🔹 Estruturas de Dados
- **Arrays e slices**
  - Criação, append, copy, slicing.
- **Maps**
  - Criação, inserção, exclusão, iteração.
- **Structs**
  - Definição, inicialização, campos exportados.
- **Exercícios**
  - Lista de compras com slice.
  - Contador de palavras usando map.
  - Struct `Pessoa` com nome e idade.

---

## 🔹 Orientação a Objetos em Go (Go way)
- **Métodos**
  - Métodos em structs.
- **Interfaces**
  - Definição, implementação implícita.
- **Composição**
  - Structs dentro de structs.
- **Exercícios**
  - Interface `Animal` com métodos `Falar()`.
  - Struct `Cachorro` e `Gato` implementando `Animal`.
  - Agenda com composição de contatos.

---

## 🔹 Concorrência e Paralelismo
- **Goroutines**
  - `go func()`.
- **Channels**
  - Comunicação entre goroutines.
  - Buffers, `select`.
- **Sync**
  - `sync.WaitGroup`, `sync.Mutex`.
- **Exercícios**
  - Programa que imprime números em paralelo.
  - Worker pool com goroutines e channels.
  - Simulação de corrida com goroutines.

---

## 🔹 Pacotes e Módulos
- **Organização**
  - `package main`, pacotes customizados.
- **Go Modules**
  - `go mod init`, `go get`.
- **Importação**
  - Pacotes padrão (`fmt`, `math`, `time`, `net/http`).
- **Exercícios**
  - Criar pacote `utils` com funções matemáticas.
  - Usar `time` para medir execução de código.

---

## 🔹 Testes e Qualidade
- **Testes**
  - `testing` package.
  - Funções `TestXxx`.
- **Benchmarks**
  - Funções `BenchmarkXxx`.
- **Lint e formatação**
  - `go fmt`, `golint`.
- **Exercícios**
  - Testar função de soma.
  - Benchmark de função de ordenação.

---

## 🔹 Web e APIs
- **HTTP básico**
  - `net/http`, servidor simples.
- **Rotas**
  - `http.HandleFunc`.
- **JSON**
  - `encoding/json`, marshal/unmarshal.
- **Frameworks**
  - `Gin`, `Echo`, `Fiber`.
- **Exercícios**
  - Servidor que retorna “Hello, World!”.
  - API CRUD de tarefas em JSON.
  - API com Gin e middleware de log.

---

## 🔹 Banco de Dados
- **SQL**
  - `database/sql`, drivers (`pq`, `mysql`).
- **ORM**
  - `GORM`.
- **Exercícios**
  - Conectar em PostgreSQL.
  - CRUD de usuários com GORM.

---

## 🔹 Projetos de Consolidação
- **Projeto 1 — CLI**
  - Ferramenta que lê arquivos CSV e gera relatórios.
- **Projeto 2 — API REST**
  - CRUD completo com Gin + GORM.
- **Projeto 3 — Concorrência**
  - Worker pool que processa imagens em paralelo.
- **Projeto 4 — Microserviço**
  - Serviço com autenticação JWT, banco de dados e Docker.

---

## 🔹 Dicas de Estudo
- Pratique diariamente com pequenos programas.
- Leia código open source em Go (ex.: Kubernetes, Hugo).
- Use `go fmt` e mantenha código idiomático.
- Prefira simplicidade: Go valoriza clareza sobre abstração.