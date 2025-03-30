package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Terms(c *gin.Context) {
	c.HTML(http.StatusOK, "private/terms", gin.H{
		"Title":   "Terms",
		"Content": "terms",
	})
}
