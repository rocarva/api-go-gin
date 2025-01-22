package database

import (
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "api-go-gin/models"
)

var DB *gorm.DB

func ConectaComBancoDeDados() {
    // String de conexão ao banco de dados
    dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
    var err error

    // Conectando ao banco de dados
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Panic("Erro ao conectar com o banco de dados:", err)
    }

    // Criar tabelas automaticamente com base nos modelos
    err = DB.AutoMigrate(&models.Aluno{})
    if err != nil {
        log.Panic("Erro ao migrar o banco de dados:", err)
    }

    log.Println("Banco de dados conectado e tabelas criadas/migradas com sucesso!")
}
