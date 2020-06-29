/**
 * @Author: youxingxiang
 * @Description:
 * @File:  wx_user
 * @Version: 1.0.0
 * @Date: 2020-06-29 16:55
 */
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp/v2"
	"os"
	"wxyapi/serializer"
)

func WxUserLogin(c *gin.Context) {
	code := c.Query("code")
	res, err := weapp.Login(os.Getenv("APPID"), os.Getenv("SECRET"), code)
	if err != nil {
		// 处理一般错误信息
		c.JSON(200, ErrorResponse(err))
	}

	if err := res.GetResponseError(); err != nil {
		// 处理微信返回错误信息
		c.JSON(200, ErrorResponse(err))
	}
	c.JSON(200, serializer.Response{Data: res})
}
