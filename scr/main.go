package main

import (
	"API_GO/models"
	"API_GO/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{ID: 1, Nome: "Setor União", Cidade: "Goiânia", Ano_Criacao: 1960},
		{ID: 2, Nome: "Vila Alpes", Cidade: "Goiânia", Ano_Criacao: 1960},
		{ID: 3, Nome: "Jardim America", Cidade: "Goiânia", Ano_Criacao: 1952},
		{ID: 4, Nome: "Novo Horizonte", Cidade: "Goiânia", Ano_Criacao: 1975},
		{ID: 5, Nome: "Jardim Europa", Cidade: "Goiânia", Ano_Criacao: 1922},
		{ID: 6, Nome: "Vila São joão", Cidade: "Goiânia", Ano_Criacao: 1968},
		{ID: 7, Nome: "Vila Nova Esperança", Cidade: "Goiânia", Ano_Criacao: 1979},
	}
	database.connectionDB()
	fmt.Println("Iniciando API em python")
	routes.HandleRequest()
}
