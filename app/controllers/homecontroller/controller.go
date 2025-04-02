package homecontroller

import "github.com/gin-gonic/gin"

// ╔═══════════════════════════════════════════════════════════╗
// ║                            PRIVATE                        ║
// ╚═══════════════════════════════════════════════════════════╝

type responseHome struct {
	Title   string
	Content string
}

type homeController struct {
	res *responseHome
}


// ╔═══════════════════════════════════════════════════════════╗
// ║                            PUBLIC                         ║
// ╚═══════════════════════════════════════════════════════════╝

type HomeController interface {
	Home(c *gin.Context)
}

func InitHomeController() HomeController {
	res := &responseHome{
		Title:  "La Belfortaine - Boucherie & Charcuterie traditionnelle à Belfort",
		Content: "home",
	}

	return &homeController{res}
}