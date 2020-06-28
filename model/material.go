/**
 * @Author: youxingxiang
 * @Description:
 * @File:  material
 * @Version: 1.0.0
 * @Date: 2020-06-28 10:55
 */
package model

import "github.com/jinzhu/gorm"

type Material struct {
	gorm.Model
	Number string `gorm:"size:32"`
	Name   string
}

func (material *Material) TableName() string {
	return "wxy_material_item"
}

func GetMaterial(Id interface{}) (Material, error) {
	var material Material
	find := DB.First(&material, Id)
	return material, find.Error
}
