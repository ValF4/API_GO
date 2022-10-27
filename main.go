package main

import (
	"API_GO/database"
	"API_GO/models"
	"API_GO/routes"
	"fmt"
)

func main() {
	models.Todosbairros = []models.AllBairros{
		{ID: 1, Nome: "Setor União", Cidade: "Goiânia", Ano_Criacao: 1960},
		{ID: 2, Nome: "Vila Alpes", Cidade: "Goiânia", Ano_Criacao: 1960},
		{ID: 3, Nome: "Jardim America", Cidade: "Goiânia", Ano_Criacao: 1952},
		{ID: 4, Nome: "Novo Horizonte", Cidade: "Goiânia", Ano_Criacao: 1975},
		{ID: 5, Nome: "Jardim Europa", Cidade: "Goiânia", Ano_Criacao: 1922},
		{ID: 6, Nome: "Vila São joão", Cidade: "Goiânia", Ano_Criacao: 1968},
		{ID: 7, Nome: "Vila Nova Esperança", Cidade: "Goiânia", Ano_Criacao: 1979},
		{ID: 8, Nome: "Cojunto Vera Cruz", Cidade: "Goiânia", Ano_Criacao: 2000},
		{ID: 9, Nome: "Setor Marista", Cidade: "Goiânia", Ano_Criacao: 1962},
		{ID: 10, Nome: "Jardim Goias", Cidade: "Goiânia", Ano_Criacao: 1980},
		{ID: 11, Nome: "Capuava", Cidade: "Goiânia", Ano_Criacao: 1950},
		{ID: 12, Nome: "Setor Negrão de Lima", Cidade: "Goiânia", Ano_Criacao: 1967},
		{ID: 13, Nome: "Setor Bueno", Cidade: "Goiânia", Ano_Criacao: 1988},
		{ID: 14, Nome: "Vila Nova", Cidade: "Goiânia", Ano_Criacao: 1963},
		{ID: 15, Nome: "Leste Universitario", Cidade: "Goiânia", Ano_Criacao: 1979},
		{ID: 16, Nome: "Setor Central", Cidade: "Goiânia", Ano_Criacao: 1933},
		{ID: 17, Nome: "Setor Aeroporto", Cidade: "Goiânia", Ano_Criacao: 1942},
		{ID: 18, Nome: "Goiânia 2", Cidade: "Goiânia", Ano_Criacao: 1997},
		{ID: 19, Nome: "Jardim Novo Mundo", Cidade: "Goiânia", Ano_Criacao: 1990},
		{ID: 20, Nome: "Jarrdim Guanabara", Cidade: "Goiânia", Ano_Criacao: 1950},
	}
	database.ConnectionDB()
	fmt.Println("Iniciando API em GO")
	routes.HandleRequest()
}
