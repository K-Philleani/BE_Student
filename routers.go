package main

import (
	"github.com/gin-gonic/gin"
	"student/apis"
)

func Routers() *gin.Engine{
	router := gin.Default()

	index := router.Group("/student")
	{
		index.POST("/insert", apis.AddStudent)
		index.GET("/getIt", apis.GetStudentOne)
		index.GET("/getAll", apis.GetStudentAll)
		index.GET("/deleteIt", apis.DeleteStudent)
		index.GET("/update", apis.UpdateStudent)
	}

	return router
}