package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectionDB() {
	strconnection := "host=localhost user=bgoiania password=postgres dbname=root port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(strconnection))
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados")
	}
}
