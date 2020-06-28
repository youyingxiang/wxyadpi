/**
 * @Author: youxingxiang
 * @Description:
 * @File:  store
 * @Version: 1.0.0
 * @Date: 2020-06-28 10:47
 */
package model

import "github.com/jinzhu/gorm"

type Store struct {
	gorm.Model
	Number     string `gorm:"size:32"`
	Name       string
	SimpleCode string
	Province   string
	City       string
	District   string
	Address    string
}

func (store *Store) TableName() string {
	return "store"
}

func GetStore(Id interface{}) (Store, error) {
	var store Store
	find := DB.First(&store, Id)
	return store, find.Error
}
