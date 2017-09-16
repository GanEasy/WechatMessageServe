package notice

import gomail "gopkg.in/gomail.v2"

//EmailNotice struct
type EmailNotice struct {
	Address string
	Title   string
	Body    string
}

//Send Notice.Run
func (n *EmailNotice) Send() {
	m := gomail.NewMessage()
	m.SetHeader("From", "GanEasy@qq.com")
	m.SetHeader("To", n.Address)
	m.SetHeader("Subject", n.Title)
	m.SetBody("text/html", n.Body)
	d := gomail.NewDialer("smtp.qq.com", 465, "GanEasy@qq.com", "yizeme2you")
	if err := d.DialAndSend(m); err != nil {
		// panic(err)
	}
}
