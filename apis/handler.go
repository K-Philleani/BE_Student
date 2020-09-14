package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"student/model"
)

// 添加学生
func AddStudent(c *gin.Context) {
	StuName := c.PostForm("stu_name")
	StuClass := c.PostForm("stu_class")
	StuSex := c.PostForm("stu_sex")
	StuAge, err := strconv.Atoi(c.PostForm("stu_age"))
	if err != nil {
		log.Println("参数转换失败, err: ", err)
		return
	}
	StuMajor := c.PostForm("stu_major")
	StuCollege := c.PostForm("stu_college")
	StuPhone := c.PostForm("stu_phone")
	StuCity := c.PostForm("stu_city")
	s := model.Student{
		StuName: StuName,
		StuClass: StuClass,
		StuSex: StuSex,
		StuAge: StuAge,
		StuMajor: StuMajor,
		StuCollege: StuCollege,
		StuPhone: StuPhone,
		StuCity: StuCity,
	}

	id, err := s.AddStudent()
	if err != nil {
		log.Println("数据库操作失败, err: ", err)
		return
	}
	msg := fmt.Sprintf("insert successful: %d", id)

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"message": msg,
		"inserted id": id,
	})
}

// 获取对应Id学生信息
func GetStudentOne(c *gin.Context) {
	StuNo, err := strconv.Atoi(c.Query("stu_no"))
	if err != nil {
		log.Println("err:", err)
		return
	}
	s := model.Student{
		StuNo: StuNo,
	}
	student, err := s.GetStudentById()
	if err != nil {
		log.Println("数据库操作失败, err: ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"message": "查询成功",
		"StuNo": StuNo,
		"student": student,
	})

}

// 获取所有学生信息
func GetStudentAll(c *gin.Context) {
	var s model.Student
	students, err := s.GetStudentAll()
	if err != nil {
		log.Println("数据查询失败, err:", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"message": "数据查询失败",
		"students": students,
	})
}

// 删除对应id学生信息
func DeleteStudent(c *gin.Context) {
	stuNo, err := strconv.Atoi(c.Query("stu_no"))
	if err != nil {
		log.Println("参数获取是失败")
		return
	}

	s := model.Student{StuNo: stuNo}
	id, err := s.DeleteStudent()
	if err != nil {
		log.Println("数据删除失败, err:", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"message": "数据删除成功",
		"stu_no": id,
	})
}
