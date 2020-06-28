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
	"strconv"
	"wxyapi/serializer"
	"wxyapi/service"
)

func StoreOrderSummary(c *gin.Context) {
	var orderService service.StoreOrderService
	if err := c.ShouldBind(&orderService); err == nil {
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
func GetOrderItemByMaterialId(c *gin.Context) {
	var orderService service.StoreOrderService
	if err := c.ShouldBind(&orderService); err == nil {
		material_id := c.Param("material_id")
		i, err := strconv.Atoi(material_id)
		if err != nil {
			c.JSON(200, ErrorResponse(err))
		}
		stores, e := orderService.GetOrderItemByMaterialId(i)
		if e != nil {
			c.JSON(200, serializer.ParamErr(e.Error(), e))
		} else {
			c.JSON(200, serializer.Response{Data: stores})
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

func StoreOrderSendItem(c *gin.Context) {
	var orderService service.StoreOrderService
	if err := c.ShouldBind(&orderService); err == nil {
		e := orderService.StoreOrderSendItems()
		if err != nil {
			c.JSON(200, serializer.ParamErr(e.Error(), e))
		} else {
			c.JSON(200, serializer.Response{})
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
