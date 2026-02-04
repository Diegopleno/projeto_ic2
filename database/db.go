package database

import (
	"fmt"
	"log"
	"os"
	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}	
	
/*
func ConectaComBancoDeDados() {
	host := os.Getenv ("DB_HOST")
	user := os.Getenv ("DB_USER")
	password := os.Getenv ("DB_PASSWORD")
	dbname := os.Getenv ("DB_NAME")
	port := os.Getenv ("DB_PORT")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)
	
	fmt.Println("DSN gerado:", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados: ", err)
	}

	DB = db
	DB.AutoMigrate(&models.Aluno{})
}
*/	
