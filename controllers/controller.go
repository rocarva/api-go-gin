package controllers

import (
	"api-go-gin/models"
	"api-go-gin/database"
	"github.com/gin-gonic/gin"
	"net/http"
	
)

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)

}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API": "E ai " + nome + ", tudo beleza ?",
	})
}


func CriaNovoALuno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao processar o JSON",
			
		})
		return 
	}
	// Incrementa o ID e adiciona o aluno ao banco de dados simulado
	database.DB.Create(&aluno)
	
	// Retorna o aluno criado 
	c.JSON(http.StatusCreated,gin.H{
		"message": "Aluno criado com sucesso",
		"aluno": aluno,
	})

}


func BuscarAlunoPorID(c *gin.Context){
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)	
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func DeletaAluno(c *gin.Context){
	id := c.Params.ByName("id")
	var aluno models.Aluno
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Aluno excluido com sucesso",
	})
}


func EditaAluno(c *gin.Context){
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao processar o JSON",
		})
		return
	}
	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, gin.H{
		"message": "Aluno alterado com sucesso",
		"aluno": aluno,
	})
}

func BuscaAlunoPorCPF(c *gin.Context){
	var aluno models.Aluno
	cpf := c.Params.ByName("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{			
			"Not found": "Aluno não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, aluno)
}
