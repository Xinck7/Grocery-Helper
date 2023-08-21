package main

import (
	"xinck/api/config"
	"xinck/api/routes"

	"gorm.io/gorm"
)

// global var
var db *gorm.DB = config.ConnectDB()

func main() {
	defer config.DisconnectDB(db)

	routes.Routes()
}
