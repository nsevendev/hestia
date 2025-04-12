package termscontroller

import "github.com/gin-gonic/gin"

// ╔═══════════════════════════════════════════════════════════╗
// ║                            PRIVATE                        ║
// ╚═══════════════════════════════════════════════════════════╝

type responseTerms struct {
	Title   string
	Content string
}

type termsController struct {
	res *responseTerms
}

// ╔═══════════════════════════════════════════════════════════╗
// ║                          PUBLIC	   	                   ║
// ╚═══════════════════════════════════════════════════════════╝

type TermsController interface {
	List(c *gin.Context)
}

func InitTermsController() TermsController {
	res := &responseTerms{
		Title:   "terms",
		Content: "terms",
	}

	return &termsController{res}
}
