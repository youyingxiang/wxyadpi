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
	"wxyapi/serializer"
	"wxyapi/service"
)

func WxUserLogin(c *gin.Context) {
	var wxUserLoginService service.WxUserLoginService
	if err := c.ShouldBind(&wxUserLoginService); err == nil {
		res := wxUserLoginService.Login()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

func WxUserDecryptUserInfo(c *gin.Context) {
	var wxUserDecryptUserInfo service.WxUserDecryptUserInfoService
	if err := c.ShouldBind(&wxUserDecryptUserInfo); err == nil {
		if xcxUser, e := GetCurrentUser(c); e == nil {
			res := wxUserDecryptUserInfo.DecryptUserInfo(xcxUser)
			c.JSON(200, res)
		} else {
			c.JSON(200, ErrorResponse(e))
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func GetWxUserInfo(c *gin.Context) {
	if xcxUser, e := GetCurrentUser(c); e != nil {
		c.JSON(200, ErrorResponse(e))
	} else {
		c.JSON(200, serializer.BuildXcxUserResponse(xcxUser))
	}
}
