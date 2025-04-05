package authcontroller

import (
	depinject "hestia/app/depInject"
	"hestia/internal/auth"

	"github.com/gin-gonic/gin"
)

// ╔═══════════════════════════════════════════════════════════╗
// ║                            PRIVATE                        ║
// ╚═══════════════════════════════════════════════════════════╝

type responseAuth struct {
	Title   string
	Content string
	Error string
}

type authController struct {
	res *responseAuth
	serviceAuth auth.AuthService
}

// ╔═══════════════════════════════════════════════════════════╗
// ║                            PUBLIC                         ║
// ╚═══════════════════════════════════════════════════════════╝

type HomeController interface {
	ShowLogin(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

func InitHomeController(c *depinject.Container) HomeController {
	res := &responseAuth{
		Title:  "Login user",
		Content: "login",
	}

	return &authController{res, c.AuthService}
}