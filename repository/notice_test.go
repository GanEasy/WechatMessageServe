// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package repository

import (
	"fmt"
	"testing"
	"time"

	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/message/custom"
)

func Test_SendText(t *testing.T) {
	SendText("o7UTkjr7if4AQgcPmveQ5wJ5alsA", "hello")
}

func Test_SendArticle(t *testing.T) {
	SendArticle("o7UTkjr7if4AQgcPmveQ5wJ5alsA", "title", "description", "", "")
}

func Test_SendNews(t *testing.T) {
	toUser := "o7UTkjr7if4AQgcPmveQ5wJ5alsA"
	// bookName := `一连十余日，都在江上行走。 123`
	// 通知用户已经关注XX小说，当该小说更新时会提醒用户查看

	ats := core.NewDefaultAccessTokenServer("wx702b93aef72f3549", "8b69f45fc737a938cbaaffc05b192394", nil)
	clt := core.NewClient(ats, nil)

	// msg := custom.NewText(toUser, bookName, "")

	news := []custom.Article{
		custom.Article{
			Title:       "为抗日牺牲的美女护士：用石头砸死日本军官，遗骸数十年才归故里",
			Description: "1938年3月，台儿庄战役爆发。中国方面守军29万人与日军5万人展开血战。战役历时一个月，中方伤亡5万余人，日方伤亡2",
			PicURL:      "http://pic3.readfollow.com/aHR0cDovL21tYml6LnFwaWMuY24vbW1iaXpfanBnL2ZZTmcxaWJPUEJ0TkhSYm5pY2tscXFJSzRnbm4waFZreTk3YTFId2oyaDRxZ3pXYmljYmNieUtwTjZlbWlhdFJSa0tXN2pFWEg0NHhVdlpXamczaEF0T2pKUS8wP3d4X2ZtdD1qcGVn",
			URL:         "http://mp.weixin.qq.com/s?__biz=MzA3MjkxMTIyNw==&mid=2650840174&idx=1&sn=93a278779739ade1ed56eff489ab3d69&chksm=84e32c7ab394a56c7182ae2beca1824d3a2a6105c31750ee2121d8914d4bd40ae0055a8879a1#rd",
		},
		custom.Article{
			Title:       "画如美人—— 读 《状态 ：艺术品的老化 》",
			Description: "假如你在卢浮宫隔着熙攘的人群望着列奥达多·达·芬奇的名画《蒙娜丽莎》出神，是否会想象或者担心这个恬静微笑着的女人慢慢变得白发苍苍、眼眶深陷、皱纹横生？",
			PicURL:      "http://pic3.readfollow.com/aHR0cDovL21tYml6LnFwaWMuY24vbW1iaXpfanBnL3BQM1AxTHlSYVJ5SWhpYmptTVkyNU9pYW5kYmFVUnJLcmlib28xR0RZWTlkVWRSVVRlQ1g2Y2liR3E1Y2o5aWNpYlV2dDVraWNFdlVJaExPb2xzZlpFZnJ5V1JsQS8wP3d4X2ZtdD1qcGVn",
			URL:         "http://mp.weixin.qq.com/s?__biz=MzA5MTkyMzgwMA==&mid=2818961058&idx=3&sn=587127aabbc6765b49c7d5967bf81393&chksm=bd85ba958af233835aa00627521a8e3a1adf64cca6bf7a3ee49b8ba5aa1f2d2c92f9cbbaa83c#rd",
		},
	}

	msg := custom.NewNews(toUser, news, "")

	e := custom.Send(clt, msg)

	if e != nil {
		fmt.Println(e.Error())
	}

	time.Sleep(time.Second)
}
