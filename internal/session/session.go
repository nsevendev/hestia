package session

import (
	"hestia/internal/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const SessionName = "hestia-session"
const SessionKeyUserID = "user_id"

func Init(storeSecret string) gin.HandlerFunc {
    store := cookie.NewStore([]byte(storeSecret))
    return sessions.Sessions(SessionName, store)
}

func SetUserSession(c *gin.Context, user *models.User) {
    session := sessions.Default(c)
    session.Set(SessionKeyUserID, user.UUID.String())
    session.Set("UserName", user.Username)
    session.Set("Email", user.Email)
    session.Save()
}

func ClearSession(c *gin.Context) {
    session := sessions.Default(c)
    session.Clear()
    session.Save()
}

func GetUserID(c *gin.Context) string {
    session := sessions.Default(c)
    id := session.Get(SessionKeyUserID)
    if idStr, ok := id.(string); ok {
        return idStr
    }
    return ""
}

func GetUserInfos(c *gin.Context) (string, string) {
    session := sessions.Default(c)
    id := session.Get(SessionKeyUserID)
    if _, ok := id.(string); ok {
        return session.Get("Email").(string), session.Get("UserName").(string)
    }
    return "", ""
}