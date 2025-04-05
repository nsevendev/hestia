package session

import (
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

func SetUserSession(c *gin.Context, userID string) {
    session := sessions.Default(c)
    session.Set(SessionKeyUserID, userID)
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