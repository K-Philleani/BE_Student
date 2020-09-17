package main

import (
	"github.com/gin-gonic/gin"
	"student/apis"
)

func Routers() *gin.Engine{
	router := gin.Default()
	account := router.Group("/account")
	{
		account.GET("/getAll", apis.GetAccountAll)
	}

	return router
}