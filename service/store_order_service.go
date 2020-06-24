/**
 * @Author: youxingxiang
 * @Description:
 * @File:  store_order
 * @Version: 1.0.0
 * @Date: 2020-06-24 09:07
 */
package service

import (
	"time"
	"wxyapi/model"
	"wxyapi/serializer"
	"wxyapi/util"
)

type StoreOrderService struct {
	Limit     int       `form:"limit"`
	Start     int       `form:"start"`
	StartTime time.Time `form:"start_time" time_format:"2006-01-02 15:04:05"`
	EndTime   time.Time `form:"end_time" time_format:"2006-01-02 15:04:05"`
}

func (service *StoreOrderService) GetStoreOrder() serializer.Response {
	order, e := model.GetStoreOrder(25)
	if e != nil {
		return serializer.ParamErr(e.Error(), e)
	}
	return serializer.Response{
		Data: order,
	}
}

func (service *StoreOrderService) GetOrderItemByMaterialId() {

}
func (service *StoreOrderService) GetStoreOrderSummary() ([]*serializer.StoreOrderSummary, error) {

	storeOrderSummarys := []*serializer.StoreOrderSummary{}
	//result := serializer.StoreOrderSummary{}
	status := util.StatusReview
	db := model.DB.Table("store_order_item").
		Select("store_order_item.material_id as material_id," +
			"count(distinct(store_order.store_id)) as store_num," +
			"sum(store_order_item.should_number) as sum_num," +
			"wxy_material_item.name as material_name,wxy_material_item.img as img," +
			"wxy_material_item.specs as material_specs," +
			"wxy_material_brand.name as material_brand").
		Joins("left join store_order on store_order_item.order_item_no = store_order.number and store_order_item.mdept_id = store_order.mdept_id " +
			"left join wxy_material_item on store_order_item.material_id = wxy_material_item.id " +
			"left join wxy_material_brand on wxy_material_item.brand_id = wxy_material_brand.id")
	if !service.StartTime.IsZero() && !service.EndTime.IsZero() {
		db = db.Where("store_order_item.created_at between ? and ?", service.StartTime, service.EndTime)
	}
	rows, e := db.Where("store_order.status= ?", status).Group("material_id").Limit(service.Limit).Offset(service.Start).Rows()

	if e != nil {
		return nil, e
	}
	defer rows.Close()
	for rows.Next() {
		storeOrderSummary := serializer.StoreOrderSummary{}
		model.DB.ScanRows(rows, &storeOrderSummary)
		storeOrderSummary.Img = util.GetQiniuImg(storeOrderSummary.Img)
		storeOrderSummarys = append(storeOrderSummarys, &storeOrderSummary)

	}
	return storeOrderSummarys, nil

}
