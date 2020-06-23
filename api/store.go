/**
 * @Author: youxingxiang
 * @Description:
 * @File:  store_order
 * @Version: 1.0.0
 * @Date: 2020-06-23 16:53
 */
package api

import "github.com/gin-gonic/gin"

type s struct {
	Data string `json:"data"`
	Msg  string `json:"msg"`
}

func StoreOrderSummary(c *gin.Context) {
	c.JSON(200, s{"sss", "还不错"})
}
