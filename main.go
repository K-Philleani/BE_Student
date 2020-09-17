package main

import (
	db "student/connection"
)

func main() {
	defer db.DB.Close()
	router := Routers()

	router.Run(":8088")
}
