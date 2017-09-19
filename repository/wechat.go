package repository

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/GanEasy/WechatMessageServe/orm"
	"github.com/yizenghui/sda/code"

	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/menu"
	"github.com/chanxuehong/wechat.v2/mp/message/callback/request"
	"github.com/chanxuehong/wechat.v2/mp/message/callback/response"
)

var (
	// 下面两个变量不一定非要作为全局变量, 根据自己的场景来选择.
	msgHandler core.Handler
	msgServer  *core.Server

//	fansCache  *cache.Cache
)

const (
	wxAppId         = "wx702b93aef72f3549"
	wxAppSecret     = "8b69f45fc737a938cbaaffc05b192394"
	wxOriId         = "gh_cb5c31e2c2dd"
	wxToken         = "admin"
	wxEncodedAESKey = ""
)

func init() {

	// conf.InitConfig("../conf/conf.toml")
	//	fansCache = cache.New(5*time.Minute, 30*time.Second)
	mux := core.NewServeMux()
	mux.DefaultMsgHandleFunc(defaultMsgHandler)
	mux.DefaultEventHandleFunc(defaultEventHandler)
	mux.MsgHandleFunc(request.MsgTypeText, textMsgHandler)
	mux.EventHandleFunc(menu.EventTypeClick, menuClickEventHandler)

	msgHandler = mux

	// fmt.Println("xx", conf.Conf.Wechat.OriID, conf.Conf.Wechat.AppID, conf.Conf.Wechat.Token, conf.Conf.Wechat.AesKey)

	msgServer = core.NewServer(wxOriId, wxAppId, wxToken, wxEncodedAESKey, msgHandler, nil)
	// msgServer = core.NewServer(conf.Conf.Wechat.OriID, conf.Conf.Wechat.AppID, conf.Conf.Wechat.Token, conf.Conf.Wechat.AesKey, msgHandler, nil)
}

func textMsgHandler(ctx *core.Context) {

	// log.Printf("收到文本消息:\n%s\n", ctx.MsgPlaintext)

	msg := request.GetText(ctx.MixedMsg)

	resp := response.NewText(msg.FromUserName, msg.ToUserName, msg.CreateTime, "请多指教")

	// ctx.AESResponse(resp, 0, "", nil) // aes密文回复

	//	resp := response.NewText(msg.FromUserName, msg.ToUserName, msg.CreateTime, msg.Content)
	ctx.RawResponse(resp) // 明文回复
	//	ctx.AESResponse(resp, 0, "", nil) // aes密文回复
}

func defaultMsgHandler(ctx *core.Context) {
	// log.Printf("收到消息:\n%s\n", ctx.MsgPlaintext)
	ctx.NoneResponse()
}

func menuClickEventHandler(ctx *core.Context) {

	event := menu.GetClickEvent(ctx.MixedMsg)

	switch key := event.EventKey; key {

	case "state":

		rc := fmt.Sprintf(`在线/离线功能开发中...`)
		resp := response.NewText(event.FromUserName, event.ToUserName, event.CreateTime, rc)
		// ctx.AESResponse(resp, 0, "", nil) // aes密文回复
		ctx.RawResponse(resp)

	default:
		resp := response.NewText(event.FromUserName, event.ToUserName, event.CreateTime, "Please look forward to more features!")
		ctx.RawResponse(resp)
		// ctx.AESResponse(resp, 0, "", nil) // aes密文回复
	}

	//ctx.RawResponse(resp) // 明文回复
	//	ctx.AESResponse(resp, 0, "", nil) // aes密文回复
}

// SetWebGetSignTaskValueForWechatPush ..
func SetWebGetSignTaskValueForWechatPush(str, openID string) bool {
	i64, err := strconv.ParseInt(str, 10, 32)
	fmt.Println(i64)
	if err != nil {
		// 这里面用正则匹配出整数
		istr := code.FindString(`(?P<int>\d+)`, str, "int")
		i64, _ = strconv.ParseInt(istr, 10, 64)
		user := orm.User{}
		user.GetUserByID(int(i64))
		if user.Registered {
			user.OpenID = openID
			user.Save()
			return true
		}
	} else {
		user := orm.User{}
		user.GetUserByID(int(i64))
		fmt.Println(user)
		if user.Registered {
			user.OpenID = openID
			user.Save()
			return true
		}
	}
	return false
}

func defaultEventHandler(ctx *core.Context) {

	event := menu.GetScanCodePushEvent(ctx.MixedMsg)

	// TODO 识别数值范围并解释要做什么事
	SetWebGetSignTaskValueForWechatPush(event.EventKey, event.FromUserName)

	rc := fmt.Sprintf(`感谢您的使用!`)
	resp := response.NewText(event.FromUserName, event.ToUserName, event.CreateTime, rc)
	// ctx.AESResponse(resp, 0, "", nil) // aes密文回复
	ctx.RawResponse(resp)
	// ctx.NoneResponse()
}

// WechatServe 微信接口服务
func WechatServe(w http.ResponseWriter, r *http.Request) {
	msgServer.ServeHTTP(w, r, nil)
}
