package main

import (
	"github.com/Shinji-Mimura/api-go-gin/database"
	"github.com/Shinji-Mimura/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
