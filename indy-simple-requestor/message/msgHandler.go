package message

import (
	//"net/http"
	"github.com/gin-gonic/gin"
)

type IssueDidRes struct {
	Did string
	Verkey string
}

func Response(c *gin.Context, status int, payload interface{}) {
	c.JSON(status, gin.H{
		"Payload" : payload,
	})
}