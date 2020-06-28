/**
 * @Author: youxingxiang
 * @Description:
 * @File:  store_order
 * @Version: 1.0.0
 * @Date: 2020-06-24 08:54
 */
package model

import "github.com/jinzhu/gorm"

type StoreOrder struct {
	gorm.Model
	Number       string
	Type         int
	StoreId      int
	Status       int
	MdepId       int
	Other        string `gorm:"size:128"`
	CreateUserId int
}

func (StoreOrder *StoreOrder) TableName() string {
	return "store_order"
}

func GetStoreOrder(Id interface{}) (StoreOrder, error) {
	var storeOrder StoreOrder
	find := DB.First(&storeOrder, Id)
	return storeOrder, find.Error
}
