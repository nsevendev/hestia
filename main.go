package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Bienvenue chez la Boucherie Hestia 🥩")
    })

    r.Run(":5000")
}