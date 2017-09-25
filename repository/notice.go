package repository

import "github.com/GanEasy/WechatMessageServe/notice"

// func init() {

// 	maxWorkers := 1
// 	maxQueue := 2
// 	//初始化一个调试者,并指定它可以操作的 工人个数
// 	dispatch := notice.NewDispatcher(maxWorkers)
// 	notice.JobQueue = make(chan notice.Job, maxQueue) //指定任务的队列长度
// 	//并让它一直接运行
// 	dispatch.Run()
// 	// close(notice.JobQueue)
// }

// SendText 发送文本
func SendText(openID, text string) (err error) {
	// openID = "o7UTkjr7if4AQgcPmveQ5wJ5alsA"
	p := notice.TextNotice{
		OpenID: openID,
		Text:   text,
	}
	notice.JobQueue <- notice.Job{
		Notice: &p,
	}
	return nil
}

// SendArticle 发送文章
func SendArticle(openID, title, description, picURL, url string) (err error) {
	p := notice.ArticleNotice{
		OpenID:      openID,
		Title:       title,
		Description: description,
		PicURL:      picURL,
		URL:         url,
	}
	notice.JobQueue <- notice.Job{
		Notice: &p,
	}
	return nil
}

// SendEmail 发送邮件
func SendEmail(address, title, body string) (err error) {
	p := notice.EmailNotice{
		Address: address,
		Title:   title,
		Body:    body,
	}
	notice.JobQueue <- notice.Job{
		Notice: &p,
	}
	return nil
}
