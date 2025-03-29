package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "public/home", gin.H{
		"title": "La Belfortaine - Boucherie & Charcuterie traditionnelle à Belfort",
		"content": "home",
	})
}