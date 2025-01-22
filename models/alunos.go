package models

import "gorm.io/gorm"
type Aluno struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Nome string `json:"nome"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}


