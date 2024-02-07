package main

import (
	"belajar-swagger/database"
	"belajar-swagger/routers"
)

func main() {
	var PORT = ":8080"
	database.StartDB()
	routers.StartServer().Run(PORT)
}
