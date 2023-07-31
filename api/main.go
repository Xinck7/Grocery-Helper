package main

import (
	"xinck/api/src/config"
	"xinck/api/src/routes"

	"gorm.io/gorm"
)

// global var
var db *gorm.DB = config.ConnectDB()

func main() {
	defer config.DisconnectDB(db)

	routes.Routes()
}
