package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/GanEasy/WechatMessageServe/orm"
	"github.com/GanEasy/WechatMessageServe/repository"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// 接入微信接口服务
func echoWxCallbackHandler(c echo.Context) error {
	repository.WechatServe(c.Response().Writer, c.Request())
	var err error
	return err
}

// Template ..
type Template struct {
	templates *template.Template
}

// Render ..
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

//Home 主页
func Home(c echo.Context) error {
	return c.Render(http.StatusOK, "home", "")
}

//Follow 关注微信公众号  	e.GET("/follow/:token", Follow)
func Follow(c echo.Context) error {
	// return c.Render(http.StatusOK, "follow", "aa")
	token := c.Param("token")
	ids := repository.Decode(token)
	if id := ids[0]; id > 0 {
		user := orm.User{}
		user.GetUserByID(id)

		url, err := repository.GetBindQrcode(int(id))
		if err != nil {
			return c.String(http.StatusOK, "-1")
		}

		// Data
		type Data struct {
			QrcodeURL string
		}

		data := &Data{
			QrcodeURL: url,
		}
		return c.Render(http.StatusOK, "follow", data)
	}
	return c.String(http.StatusOK, "0")
}

//Reg 注册
func Reg(c echo.Context) error {

	host := "https://readfollow.com"
	// host = "http://192.168.1.152:8888"
	email := c.Param("email")
	user := orm.User{}
	user.GetUserByEmail(email)
	if !user.Registered && user.Subscribed { // 未注册，已订阅

		title := "欢迎加入 voous "
		// title := "Welcome to join voous"
		token := repository.Encode([]int{int(user.ID), 1})
		// url, err := repository.GetBindQrcode(int(user.ID))
		// if err != nil {
		// 	return c.String(http.StatusOK, "-1")
		// }
		// body := fmt.Sprintf(
		// 	`Hello , Please click the link, <a href="%v/join/%v">to join us</a> !

		// 	<br>
		// 	<img src="%v" />
		// 	<br>
		// 	please use	WeChat scans in 30 minutes, If you need notification

		// 	<br><br> <a href="%v/unsubscribed/%v">I don't want to accept any more mail</a>`, host, token, url, host, token)
		body := fmt.Sprintf(
			`您好 , 请点击链接, <a href="%v/join/%v">加入我们</a> !

				
				<br>

				我们提供免费的微信消息提醒服务，如需帮助，请直接回复邮件！

				<br><br> <a href="%v/unsubscribed/%v">我不希望再收到任何邮件</a>`, host, token, host, token)
		repository.SendEmail(email, title, body)

		user.Invited = true
		user.Save()
		return c.String(http.StatusOK, "1")

	}
	return c.String(http.StatusOK, "0")
}

//Join 加入
func Join(c echo.Context) error {
	token := c.Param("token")
	ids := repository.Decode(token)
	if id := ids[0]; id > 0 {
		user := orm.User{}
		user.GetUserByID(id)
		if user.Email != "" && user.Registered == false {
			user.Registered = true
			user.Save()
			return c.Redirect(http.StatusFound, fmt.Sprintf("/follow/%v", token))
		}
	}
	return c.Redirect(http.StatusFound, "/")
}

//Unsubscribed 不再关注(不要再向用户发送任何邮件)
func Unsubscribed(c echo.Context) error {
	token := c.Param("token")
	ids := repository.Decode(token)
	if id := ids[0]; id > 0 {
		user := orm.User{}
		user.GetUserByID(id)
		if user.Email != "" && user.Subscribed == true {
			user.Subscribed = false
			user.Save()
			return c.String(http.StatusOK, "系统将不会再向您发送任何邮件！")
		}
	}
	return c.String(http.StatusOK, "出现一些小问题，但想必您的期盼已满足！")
}

//Text 文本
func Text(c echo.Context) error {
	email := c.Param("email")
	user := orm.User{}
	user.GetUserByEmail(email)
	if user.Invited && user.Registered && user.Subscribed && user.OpenID != "" {
		str := c.QueryParam("s")
		if str != "" {
			repository.SendText(user.OpenID, str)
			return c.String(http.StatusOK, "1")
		}
	}
	return c.String(http.StatusOK, "0")
}

//Article 文章
func Article(c echo.Context) error {
	email := c.Param("email")
	user := orm.User{}
	user.GetUserByEmail(email)
	if user.Invited && user.Registered && user.Subscribed && user.OpenID != "" {
		title := c.QueryParam("title")
		description := c.QueryParam("description")
		cover := c.QueryParam("cover")
		url := c.QueryParam("url")
		if title != "" && description != "" {
			repository.SendArticle(user.OpenID, title, description, cover, url)
			return c.String(http.StatusOK, "1")
		}
	}
	return c.String(http.StatusOK, "0")
}
func main() {
	orm.DB().AutoMigrate(&orm.User{})

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	e.Pre(middleware.HTTPSRedirect())
	e.Pre(middleware.HTTPSNonWWWRedirect())
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", Home)

	e.GET("/reg/:email", Reg)

	e.GET("/join/:token", Join)

	e.GET("/follow/:token", Follow)

	e.GET("/unsubscribed/:token", Unsubscribed)

	e.GET("/text/:email", Text)

	e.GET("/article/:email", Article)

	e.File("/favicon.ico", "images/favicon.ico")

	e.Any("/wx_callback", echoWxCallbackHandler)

	// e.Static("/", "src")
	// Start server
	// e.Logger.Fatal(e.Start(":8888"))
	// go e.Logger.Fatal(e.Start(":80"))
	e.Logger.Fatal(e.StartAutoTLS(":443"))

}
