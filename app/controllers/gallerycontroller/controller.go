package gallerycontroller

import "github.com/gin-gonic/gin"

// ╔═══════════════════════════════════════════════════════════╗
// ║                            PRIVATE                        ║
// ╚═══════════════════════════════════════════════════════════╝

type responseGallery struct {
	Title   string
	Content string
}

type galleryController struct {
	res *responseGallery
}

// ╔═══════════════════════════════════════════════════════════╗
// ║                          PUBLIC	   	                   ║
// ╚═══════════════════════════════════════════════════════════╝

type GalleryController interface {
	First(c *gin.Context)
}

func InitGalleryController() GalleryController {
	res := &responseGallery{
		Title:   "gallery",
		Content: "gallery",
	}

	return &galleryController{res}
}