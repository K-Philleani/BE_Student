package main

import (
	"github.com/gin-gonic/gin"
	"student/apis"
)

func Routers() *gin.Engine{
	router := gin.Default()

	router.POST("/student/insert", apis.AddStudent)
	router.GET("/student/getIt", apis.GetStudentOne)
	router.GET("/student/getAll", apis.GetStudentAll)
	router.GET("/student/deleteIt", apis.DeleteStudent)
	return router
}