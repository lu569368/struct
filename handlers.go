package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/student/gingo/zuoye/dbops"
)

//GetBook ...
func GetBook(c *gin.Context) {
	books, err := dbops.GetAllbook()
	if err != nil {
		fmt.Println("获取参数失败", err)
	}
	for i, ele := range books {
		fmt.Printf("comment: %d, %v \n", i, ele)
	}
	// // c.String(200,books)
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "我是测试", "ce": books})
}
