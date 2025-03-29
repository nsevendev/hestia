package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "private/dashboard", gin.H{
		"title": "Dashboard",
		"content": "dashboard",
	})
}