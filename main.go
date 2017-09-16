package main

import (
	"fmt"
	"net/http"

	"github.com/GanEasy/WechatMessageServe/orm"
	"github.com/GanEasy/WechatMessageServe/repository"
	"github.com/labstack/echo"
)

// 接入微信接口服务
func echoWxCallbackHandler(c echo.Context) error {
	repository.WechatServe(c.Response().Writer, c.Request())
	var err error
	return err
}

func main() {

	host := "http://192.168.1.152:8888"
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	e.GET("/reg/:email", func(c echo.Context) error {
		email := c.Param("email")
		user := orm.User{}
		user.GetUserByEmail(email)
		if !user.Registered && user.Subscribed { // 未注册，已订阅

			// title := "欢迎加入 voous "
			title := "Welcome to join voous"
			token := repository.Encode([]int{int(user.ID), 1})
			url, err := repository.GetBindQrcode(int(user.ID))
			if err != nil {
				return c.String(http.StatusOK, "-1")
			}
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
					<img src="%v" />
					<br>
					如果您需要使用我们的微信通知，请使用微信扫描二维码 (需要关注公众号)
					<br>

					我们提供免费的微信消息提醒服务，如需帮助，请直接回复邮件！

					<br><br> <a href="%v/unsubscribed/%v">我不希望再收到任何邮件</a>`, host, token, url, host, token)
			repository.SendEmail(email, title, body)

			user.Invited = true
			user.Save()
			return c.String(http.StatusOK, "1")

		}
		return c.String(http.StatusOK, "0")
	})

	e.GET("/join/:token", func(c echo.Context) error {
		token := c.Param("token")
		ids := repository.Decode(token)
		if id := ids[0]; id > 0 {
			user := orm.User{}
			user.GetUserByID(id)
			if user.Email != "" && user.Registered == false {
				user.Registered = true
				user.Save()
				return c.String(http.StatusOK, "1")
			}
		}
		return c.String(http.StatusOK, "0")
	})

	e.GET("/unsubscribed/:token", func(c echo.Context) error {
		token := c.Param("token")
		ids := repository.Decode(token)
		if id := ids[0]; id > 0 {
			user := orm.User{}
			user.GetUserByID(id)
			if user.Email != "" && user.Subscribed == true {
				user.Subscribed = false
				user.Save()
				return c.String(http.StatusOK, "1")
			}
		}
		return c.String(http.StatusOK, "0")
	})

	e.GET("/text/:email", func(c echo.Context) error {
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
	})

	e.File("/favicon.ico", "images/favicon.ico")

	e.Any("/wx_callback", echoWxCallbackHandler)

	// e.Static("/", "src")
	// Start server
	e.Logger.Fatal(e.Start(":8888"))
}
