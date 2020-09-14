package main

import (
	db "student/connection"
)

func main() {
	defer db.SqlDB.Close()
	router := Routers()

	router.Run(":9999")
}
