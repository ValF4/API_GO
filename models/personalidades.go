package models

type Personalidade struct {
	ID          int    `json:"ID"`
	Nome        string `json:"Nome"`
	Cidade      string `json:"Cidade"`
	Historia    string `json:"Historia"`
	Ano_Criacao int16  `json:"Ano_Criacao"`
}

var Personalidades []Personalidade
