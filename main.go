package main

import (
	"github.com/rsenoi/CIGO2/database"
	"github.com/rsenoi/CIGO2/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
