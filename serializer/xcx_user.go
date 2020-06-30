/**
 * @Author: youxingxiang
 * @Description:
 * @File:  xcx_user
 * @Version: 1.0.0
 * @Date: 2020-06-30 14:05
 */
package serializer

import "wxyapi/model"

type XcxUser struct {
	ID         uint   `json:"id"`
	Nickname   string `json:"nickname"`
	Sex        int    `json:"sex"`
	Avatar     string `json:"avatar"`
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	Province   string `json:"province"`
	City       string `json:"city"`
	Phone      string `json:"phone"`
	CreatedAt  int64  `json:"created_at"`
}

// BuildUser 序列化用户
func BuildXcxUser(user *model.XcxUser) XcxUser {
	return XcxUser{
		ID:        user.ID,
		Openid:    user.Openid,
		Unionid:   user.Unionid,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		Sex:       user.Sex,
		Province:  user.Province,
		City:      user.City,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

// BuildUserResponse 序列化用户响应
func BuildXcxUserResponse(user *model.XcxUser) Response {
	return Response{
		Data: BuildXcxUser(user),
	}
}
