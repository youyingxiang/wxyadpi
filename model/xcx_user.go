/**
 * @Author: youxingxiang
 * @Description:
 * @File:  admin_user
 * @Version: 1.0.0
 * @Date: 2020-06-30 10:22
 */
package model

import "github.com/jinzhu/gorm"

type XcxUser struct {
	gorm.Model
	Nickname   string `json:"nickname"`
	Sex        int    `json:"sex"`
	Avatar     string `json:"avatar"`
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	Province   string `json:"province"`
	City       string `json:"city"`
	Phone      string `json:"phone"`
}

func GetXcxUser(ID interface{}) (xcx_user XcxUser, err error) {
	err = DB.First(&xcx_user, ID).Error
	return
}
