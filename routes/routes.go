package routes

import (
	"api-go-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandRequest() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome/", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoALuno)
	r.GET("/alunos/:id", controllers.BuscarAlunoPorID)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.Run()
}
