/**
 * @Author: youxingxiang
 * @Description:
 * @File:  wx_user_service
 * @Version: 1.0.0
 * @Date: 2020-06-30 09:44
 */
package service

import (
	"github.com/jinzhu/gorm"
	"github.com/medivhzhan/weapp/v2"
	"os"
	"wxyapi/model"
	"wxyapi/serializer"
)

type WxUserLoginService struct {
	Code string `json:"code" form:"code" binding:"required"`
}

func (service *WxUserLoginService) Login() serializer.Response {
	response, err := weapp.Login(os.Getenv("APPID"), os.Getenv("SECRET"), service.Code)

	if err != nil {
		// 处理一般错误信息
		return serializer.ParamErr(err.Error(), err)
	}
	if err := response.GetResponseError(); err != nil {
		// 处理微信返回错误信息
		return serializer.ParamErr(err.Error(), err)
	}
	if err := service.login(response); err != nil {
		return serializer.ParamErr(err.Error(), err)
	}

	return serializer.Response{Data: response}
}

func (service *WxUserLoginService) login(response *weapp.LoginResponse) (err error) {
	user := model.XcxUser{}
	err = model.DB.Where(model.XcxUser{Openid: response.OpenID}).First(&user).Error
	if err != nil {
		// 未找到数据
		if gorm.IsRecordNotFoundError(err) {
			user.Openid = response.OpenID
			user.SessionKey = response.SessionKey
			model.DB.Create(user)
			return nil
		}
		return
	}
	user.SessionKey = response.SessionKey
	model.DB.Save(user)
	return
}

type WxUserDecryptUserInfoService struct {
	EncryptedData string `form:"encrypted_data" binding:"required"`
	RawData       string `form:"raw_data" binding:"required"`
	Signature     string `form:"signature" binding:"required"`
	Iv            string `form:"iv" binding:"required"`
	//OpenId        string `form:"openid" binding:"required"`
}

func (service *WxUserDecryptUserInfoService) DecryptUserInfo(user *model.XcxUser) serializer.Response {
	if e := service.decryptUserInfo(user); e != nil {
		return serializer.ParamErr(e.Error(), e)
	}
	return serializer.Response{Data: user}

}

func (service *WxUserDecryptUserInfoService) decryptUserInfo(user *model.XcxUser) error {
	info, err := weapp.DecryptUserInfo(user.SessionKey, service.RawData, service.EncryptedData, service.Signature, service.Iv)
	if err != nil {
		return err
	}
	user.Avatar = info.Avatar
	user.Nickname = info.Nickname
	user.City = info.City
	user.Province = info.Province
	user.Sex = info.Gender
	user.Unionid = info.UnionID
	model.DB.Save(user)
	return nil
}
