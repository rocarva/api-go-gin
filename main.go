package main

import (
	"api-go-gin/routes"
	"api-go-gin/database"
)

func main() {
	database.ConectaComBancoDeDados()	
	routes.HandRequest()
}
