package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "public/home", gin.H{
		"Title": "La Belfortaine - Boucherie & Charcuterie traditionnelle à Belfort",
		"Content": "home",
	})
}