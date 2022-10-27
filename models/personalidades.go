package models

type Personalidade struct {
	ID          int    `json:"ID"`
	Nome        string `json:"Nome"`
	Cidade      string `json:"Cidade"`
	Ano_Criacao int16  `json:"Ano_Criacao"`
}

var Personalidades []Personalidade
