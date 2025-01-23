package database

import (
	"api-go-gin/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConectaComBancoDeDados() {
	// Caminho do arquivo SQLite
	dsn := "alunos.db" // O banco de dados será salvo no arquivo alunos.db
	var err error

	// Conexão com o banco de dados SQLite
	DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados SQLite:", err)
	}

	// Migração automática dos modelos
	err = DB.AutoMigrate(&models.Aluno{})
	if err != nil {
		log.Panic("Erro ao migrar o banco de dados:", err)
	}

	log.Println("Banco de dados SQLite conectado e tabelas criadas/migradas com sucesso!")
}
