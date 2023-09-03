package main

import (
	"github.com/DeniSaputra55/task-5-vix-btpns-DeniSaputra/database"
	"github.com/DeniSaputra55/task-5-vix-btpns-DeniSaputra/router"
)

func main() {
	database.InitDB()
	database.MigrateDB()
	r := router.RouteInit()
	r.Run(":9000")
}
