package middleware

import (
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	DBMongoPool = "db_mongo_pool"
	DBRedisPool = "redis_pool"
)

type Middleware struct{}

func (m Middleware) HandleAuthLevel(auth int, endpoint gin.HandlerFunc) []gin.HandlerFunc {
	var rtn []gin.HandlerFunc
	switch auth {
	case 0: // grant access to everyone
	case 1: // check for session
		rtn = append(rtn, m.VerifySession)
	}
	rtn = append(rtn, endpoint)

	return rtn
}

func (m Middleware) Request(c *gin.Context) {
	_, err := httputil.DumpRequest(c.Request, true)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to parse request"})
		return
	}

	// prevent log binary file, eg. upload image file
	ct := c.Request.Header.Get("Content-Type")
	if strings.HasPrefix(ct, "application/json") {
	}

	c.Next()
}

func (m Middleware) VerifySession(c *gin.Context) {
	// assume check session from redis
	session := false

	if !session {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "you don't have authorize to call this method."})
		return
	}
	c.Next()
}
