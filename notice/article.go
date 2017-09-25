package notice

import (
	"fmt"
	"time"

	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/message/custom"
)

//ArticleNotice struct
type ArticleNotice struct {
	OpenID      string
	Title       string
	Description string
	PicURL      string
	URL         string
}

//Send Notice.Run
func (n *ArticleNotice) Send() {

	toUser := "o7UTkjr7if4AQgcPmveQ5wJ5alsA"
	// bookName := `一连十余日，都在江上行走。 123`
	// 通知用户已经关注XX小说，当该小说更新时会提醒用户查看

	ats := core.NewDefaultAccessTokenServer("wx702b93aef72f3549", "8b69f45fc737a938cbaaffc05b192394", nil)
	clt := core.NewClient(ats, nil)

	// msg := custom.NewText(toUser, bookName, "")

	news := []custom.Article{
		custom.Article{
			Title:       n.Title,
			Description: n.Description,
			PicURL:      n.PicURL,
			URL:         n.URL,
		},
	}

	msg := custom.NewNews(toUser, news, "")

	e := custom.Send(clt, msg)

	if e != nil {
		fmt.Println(e.Error())
	}

	time.Sleep(time.Second)
}
