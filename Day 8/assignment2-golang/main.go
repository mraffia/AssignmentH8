package main

import (
	"assignment2-golang/database"
	"assignment2-golang/routers"
)

func main() {
	var PORT = ":8080"
	database.StartDB()
	routers.StartServer().Run(PORT)
}
