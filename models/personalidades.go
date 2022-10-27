package models

type AllBairros struct {
	ID          int    `json:"ID"`
	Nome        string `json:"Nome"`
	Cidade      string `json:"Cidade"`
	Ano_Criacao int16  `json:"Ano_Criacao"`
}

var Todosbairros []AllBairros
