package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"wxyapi/model"
	"wxyapi/serializer"
	"wxyapi/util"
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
			user := model.XcxUser{}
			err := model.DB.Where(model.XcxUser{Openid: openid}).First(&user).Error
			if err != nil {
				if gorm.IsRecordNotFoundError(err) {
					c.JSON(200, serializer.Response{})
				} else {
					c.JSON(200, serializer.ParamErr(err.Error(), err))
				}

				c.Abort()
			}
			c.Set(util.CTX_XCX_USER, &user)
			c.Next()
			return
		}
		c.JSON(200, serializer.NotAuth())
		c.Abort()
	}
}
