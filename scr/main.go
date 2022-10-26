package main

import (
	"API_GO/models"
	"API_GO/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{ID: 1, Nome: "Setor União", Cidade: "Goiânia", Historia: "historia 1", Ano_Criacao: 1960},
		{ID: 2, Nome: "Vila Alpes", Cidade: "Goiânia", Historia: "historia 2", Ano_Criacao: 1960},
		{ID: 3, Nome: "Jardim America", Cidade: "Goiânia", Historia: "historia 3", Ano_Criacao: 1952},
		{ID: 4, Nome: "Novo Horizonte", Cidade: "Goiânia", Historia: "historia 4", Ano_Criacao: 1975},
		{ID: 5, Nome: "Jardim Europa", Cidade: "Goiânia", Historia: "historia 5", Ano_Criacao: 1922},
		{ID: 6, Nome: "Vila São joão", Cidade: "Goiânia", Historia: "historia 6", Ano_Criacao: 1968},
		{ID: 7, Nome: "Vila Nova Esperança", Cidade: "Goiânia", Historia: "historia 7", Ano_Criacao: 1979},
	}
	fmt.Println("Iniciando API em python")
	routes.HandleRequest()
}
