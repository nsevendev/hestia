package gallerycontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (gc *galleryController)  First(c *gin.Context) {
	c.HTML(http.StatusOK, "private/gallery", gc.res)
}
