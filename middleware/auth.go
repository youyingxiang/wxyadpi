package middleware

import (
	"wxyapi/model"
	"wxyapi/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("user_id")
		if uid != nil {
			user, err := model.GetUser(uid)
			if err == nil {
				c.Set("user", &user)
			}
		}
		c.Next()
	}
}

// AuthRequired 需要登录
func AuthRequired() func(c *gin.Context) {
	return func(c *gin.Context) {
		if openid := c.GetHeader("openid"); len(openid) > 0 {
			c.Next()
			return
		}
		c.JSON(200, serializer.NotAuth())
		c.Abort()
	}
}
