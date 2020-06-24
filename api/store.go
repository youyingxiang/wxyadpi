/**
 * @Author: youxingxiang
 * @Description:
 * @File:  store_order
 * @Version: 1.0.0
 * @Date: 2020-06-23 16:53
 */
package api

import (
	"github.com/gin-gonic/gin"
	"wxyapi/serializer"
	"wxyapi/service"
)

func StoreOrderSummary(c *gin.Context) {
	orderService := &service.StoreOrderService{}
	if err := c.ShouldBind(orderService); err == nil {
		summaries, e := orderService.GetStoreOrderSummary()

		if e != nil {
			c.JSON(200, serializer.ParamErr(e.Error(), e))
		} else {
			c.JSON(200, serializer.Response{Data: summaries})
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}
