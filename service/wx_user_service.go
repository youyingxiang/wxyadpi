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
	Code          string `json:"code" form:"code" binding:"required"`
	EncryptedData string `form:"encrypted_data" binding:"required"`
	//RawData       string `form:"raw_data" binding:"required"`
	//Signature     string `form:"signature" binding:"required"`
	Iv string `form:"iv" binding:"required"`
}

func (service *WxUserLoginService) Login() serializer.Response {
	// 调用微信登陆
	loginResponse, err := service.login()

	if err != nil {
		return serializer.ParamErr(err.Error(), err)
	}
	// openid保存数据库
	user, err := service.saveOpenid(loginResponse)

	if err != nil {
		return serializer.ParamErr(err.Error(), err)
	}
	// 解密
	return serializer.Response{Data: user}
}

func (service *WxUserLoginService) login() (response *weapp.LoginResponse, err error) {
	response, err = weapp.Login(os.Getenv("APPID"), os.Getenv("SECRET"), service.Code)
	if err != nil {
		return
	}
	if err = response.GetResponseError(); err != nil {
		return nil, err
	}
	return
}

func (service *WxUserLoginService) saveOpenid(response *weapp.LoginResponse) (*model.XcxUser, error) {
	user := &model.XcxUser{}
	// 解密加密的用户信息
	if info, err := service.decryptUserInfo(response); err != nil {
		return nil, err
	} else {
		if err := model.DB.Where(model.XcxUser{Openid: response.OpenID}).First(user).Error; err != nil {
			// 未找到数据
			if gorm.IsRecordNotFoundError(err) {
				saveUser := service.buildSaveUser(response, info, user)
				model.DB.Create(saveUser)
				return saveUser, nil
			}
			return nil, err
		}
		saveUser := service.buildSaveUser(response, info, user)
		model.DB.Save(saveUser)
		return saveUser, nil
	}
}

func (service *WxUserLoginService) buildSaveUser(response *weapp.LoginResponse, info *WxUserInfo, user *model.XcxUser) *model.XcxUser {
	user.Openid = response.OpenID
	//user.SessionKey = response.SessionKey
	user.Avatar = info.Avatar
	user.Nickname = info.Nickname
	user.City = info.City
	user.Province = info.Province
	user.Sex = info.Gender
	user.Unionid = info.UnionID
	return user
}

func (service *WxUserLoginService) decryptUserInfo(response *weapp.LoginResponse) (info *WxUserInfo, err error) {
	crypt := NewWXUserDataCrypt(os.Getenv("APPID"), response.SessionKey)
	info, err = crypt.Decrypt(service.EncryptedData, service.Iv)
	//info, err = weapp.DecryptUserInfo(response.SessionKey, "", service.EncryptedData, service.Signature, service.Iv)
	return
}

type WxUserDecryptUserInfoService struct {
	EncryptedData string `form:"encrypted_data" binding:"required"`
	Iv            string `form:"iv" binding:"required"`
}

func (service *WxUserDecryptUserInfoService) DecryptUserInfo(user *model.XcxUser) serializer.Response {
	if e := service.decryptUserInfo(user); e != nil {
		return serializer.ParamErr(e.Error(), e)
	}
	return serializer.Response{Data: user}
}

func (service *WxUserDecryptUserInfoService) decryptUserInfo(user *model.XcxUser) error {
	//info, err := weapp.DecryptUserInfo(user.SessionKey, service.RawData, service.EncryptedData, service.Signature, service.Iv)
	crypt := NewWXUserDataCrypt(os.Getenv("APPID"), user.SessionKey)
	info, err := crypt.Decrypt(service.EncryptedData, service.Iv)
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
