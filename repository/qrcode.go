package repository

import (
	"fmt"

	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/qrcode"
)

//CreateTempQrcode 创建临时二维码
func CreateTempQrcode(id int32) (*qrcode.TempQrcode, error) {
	ats := core.NewDefaultAccessTokenServer("wx702b93aef72f3549", "8b69f45fc737a938cbaaffc05b192394", nil)
	clt := core.NewClient(ats, nil)
	return qrcode.CreateTempQrcode(clt, id, 7200)
}

// GetBindQrcode 站点获取签名任务
func GetBindQrcode(uid int) (url string, err error) {
	qrcode, err := CreateTempQrcode(int32(uid))
	if err != nil {
		return "", err
	}
	url = fmt.Sprintf("https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=%v", qrcode.Ticket)
	return url, nil
}
