package gallerycontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (gc *galleryController)  First(c *gin.Context) {
	galleryFirst, err := gc.galleryService.GetFirst()
	if err != nil {
		c.HTML(http.StatusOK, "private/gallery", gc.res)
	}

	gc.res.GalleryFirst = galleryFirst

	c.HTML(http.StatusOK, "private/gallery", gc.res)
}
