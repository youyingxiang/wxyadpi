package util

import (
	"math/rand"
	"os"
	"strings"
	"time"
)

// RandStringRunes 返回随机字符串
func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GetQiniuImg(url string) string {
	if url == "" {
		url = os.Getenv("DEFAULT_IMG")
	}
	qiniuUrl := os.Getenv("QINIU_URL")
	return strings.Join([]string{qiniuUrl, url}, "/")
}

type Status int

const (
	StatusWaitReview Status = iota // value --> 0
	StatusReview                   // value --> 1
	StatusReReview                 // value --> 2
	StatusExcute                   // value --> 3
	StatusFinish
)

const (
	CTX_XCX_USER = "xcx_user"
)

type StoreItemStatus int

const (
	StatusNo StoreItemStatus = iota
	StatusOk
)

//func (this Status) String() string {
//	switch this {
//	case StatusWaitReview:
//		return "待审核"
//	case StatusReview:
//		return "审核"
//	case StatusReReview:
//		return "反审核"
//	case StatusExcute:
//		return "已发货"
//	case StatusFinish:
//		return "已完成"
//	default:
//		return "Unknow"
//	}
//}
