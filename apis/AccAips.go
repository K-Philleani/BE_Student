package apis

import (
	"github.com/gin-gonic/gin"
	"student/model"
)

func GetAccountAll(c *gin.Context) {
	var account model.Account
	accounts, err := account.GetAccountAll()
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"message": "查询失败",
		})
		panic(err)
	}
	c.JSON(200, gin.H{
		"code": 1,
		"message": "查询成功",
		"accountList": accounts,
	})
}