package gallerycontroller

import (
	depinject "hestia/app/depInject"
	"hestia/internal/gallery"
	"hestia/internal/models"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ╔═══════════════════════════════════════════════════════════╗
// ║                            PRIVATE                        ║
// ╚═══════════════════════════════════════════════════════════╝

type responseGallery struct {
	Title   string
	Content string
	GalleryFirst *models.Gallery
	Error   string
	Success string
}

type galleryController struct {
	res *responseGallery
	galleryService gallery.GalleryService
}

func validateDataStringEmpty(c *gin.Context , value *string, name string) string {
	if *value == "" {
		c.Redirect(
			http.StatusSeeOther, 
			"/dashboard/gallery?statusCode=" + strconv.Itoa(http.StatusBadRequest) + "&error=" + url.QueryEscape("Erreur de validation, " + name + " requis"),
		)
	}

	return  *value
}

// ╔═══════════════════════════════════════════════════════════╗
// ║                          PUBLIC	   	                   ║
// ╚═══════════════════════════════════════════════════════════╝

type GalleryController interface {
	First(c *gin.Context)
	AddImage(c *gin.Context)
	DeleteImageById(c *gin.Context)
}

func InitGalleryController(c *depinject.Container) GalleryController {
	res := &responseGallery{
		Title:   "gallery",
		Content: "gallery",
	}

	return &galleryController{res, c.GalleryService}
}