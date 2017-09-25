package notice

import (
	"fmt"

	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/message/custom"
	"github.com/chanxuehong/wechat.v2/mp/message/template"
)

//TextNotice struct
type TextNotice struct {
	OpenID string
	Text   string
}

// FollowMSG 关注通知消息结构
type FollowMSG struct {
	Name template.DataItem `json:"name"`
	// Name template.DataItem `json:"content"`
}

//Send Notice.Run
func (n *TextNotice) Send() {

	fmt.Printf("run text send")
	ats := core.NewDefaultAccessTokenServer("wx702b93aef72f3549", "8b69f45fc737a938cbaaffc05b192394", nil)
	clt := core.NewClient(ats, nil)

	msg := custom.NewText(n.OpenID, n.Text, "")

	e := custom.Send(clt, msg)

	if e != nil {

		fmt.Println(e.Error())

		msg := template.TemplateMessage2{
			ToUser:     n.OpenID,
			TemplateId: "jN1NL9EJvG2bD8EtejiH6liVcxynfHhcl-AlwEwM-l0",
			// TemplateId: "EysP-Y4edh8ZoYDco58ULJESKXSk_n7ulKCJ4UCzeoo", // 测试
			URL: "",
			Data: FollowMSG{
				Name: template.DataItem{Value: n.Text, Color: ""},
			},
		}

		template.Send(clt, msg)
	}

	fmt.Printf("text send")
}
