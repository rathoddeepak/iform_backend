package middleware;

import (
	sessionModel "iform/internal/models/session"
	"github.com/gin-gonic/gin"
	"net/http"
);

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Session, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func UserFromSessionID() gin.HandlerFunc {
  return func(c *gin.Context) {
    sessionId := c.GetHeader("X-Session")
    if sessionId == "" {
      c.AbortWithStatus(http.StatusUnauthorized)
      return
    }
    sessionData, err := sessionModel.GetSessionData(sessionId)
    if err != nil {
      c.AbortWithStatus(http.StatusInternalServerError)
      return
    }
    if sessionData.UserId == 0 {
      c.AbortWithStatus(http.StatusUnauthorized)
      return
    }
    c.Set("user_id", sessionData.UserId)
    c.Next()
  }
}

func IfUserFromSessionID() gin.HandlerFunc {
  return func(c *gin.Context) {
    sessionId := c.GetHeader("X-Session")
    if sessionId != "" {
      sessionData, _ := sessionModel.GetSessionData(sessionId)    
      c.Set("user_id", sessionData.UserId)      
    }
    c.Next();
  }
}
