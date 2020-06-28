/**
 * @Author: youxingxiang
 * @Description:
 * @File:  store_order_item
 * @Version: 1.0.0
 * @Date: 2020-06-28 12:09
 */
package model

import "github.com/jinzhu/gorm"

type StoreOrderItem struct {
	gorm.Model
	OrderItemNo  string `gorm:"size:32"`
	MaterialId   int
	Status       int
	Price        float64
	ShouldNumber int
	ActualNumber int
	MdepId       int
	StoreOrder   StoreOrder `gorm:"foreignkey:OrderItemNo;association_foreignkey:Number"`
}

func (storeOrderItem *StoreOrderItem) TableName() string {
	return "store_order_item"
}

func GetStoreOrderItem(Id interface{}) (StoreOrderItem, error) {
	var storeOrderItem StoreOrderItem
	e := DB.Where("id = ?", Id).Preload("StoreOrder").First(&storeOrderItem).Error
	return storeOrderItem, e
}
