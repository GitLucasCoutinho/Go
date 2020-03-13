package main

import (
	"encoding/json" // pacote para trabalhar com JSON (serialização e desserialização)
	"fmt"           // pacote para imprimir no console
	"io/ioutil"     // pacote para ler dados de streams (como resposta HTTP)
	"net/http"      // pacote para fazer requisições HTTP
)

// Estrutura que representa o formato dos dados retornados pela API do IBGE
// Cada município tem um ID e um Nome
type Municipio struct {
	ID   int    `json:"id"`   // campo "id" do JSON será mapeado para ID
	Nome string `json:"nome"` // campo "nome" do JSON será mapeado para Nome
}

func main() {
	// URL da API pública do IBGE (contem lista de municípios do estado de São Paulo - código 35)
	url := "https://servicodados.ibge.gov.br/api/v1/localidades/estados/35/municipios"

	// Fazendo requisição HTTP GET para a API
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erro ao acessar API:", err) // caso ocorra erro na requisição
		return
	}
	defer resp.Body.Close() // garante que o corpo da resposta será fechado ao final

	// Lendo todo o conteúdo da resposta HTTP
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler resposta:", err) // caso ocorra erro na leitura
		return
	}

	// Imprimindo o JSON bruto retornado pela API
	fmt.Println("JSON bruto da API do IBGE:")
	fmt.Println(string(body))

	// Convertendo (desserializando) o JSON para slice de structs Municipio
	var municipios []Municipio
	err = json.Unmarshal(body, &municipios)
	if err != nil {
		fmt.Println("Erro ao converter JSON:", err) // caso ocorra erro na conversão
		return
	}

	// Exibindo alguns dados tratados (os 10 primeiros municípios)
	fmt.Println("\nLista de municípios de São Paulo:")
	for i, m := range municipios {
		if i < 10 { // limita a impressão aos 10 primeiros
			fmt.Printf("ID: %d - Nome: %s\n", m.ID, m.Nome)
		}
	}
}