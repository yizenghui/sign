package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	c "github.com/yizenghui/sign/api"
	cpi "github.com/yizenghui/sign/core"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "welcome to fire minapp api, this build by  go!")
	})

	// 获取用户签名 token id
	e.GET("/gettoken", c.GetToken)

	// 获取推荐码(图片资源)
	e.GET("/qrcode", func(c echo.Context) error {
		scene := c.QueryParam("scene")
		page := `pages/index/index`
		if scene == "" {
			return c.HTML(http.StatusOK, "")
		}
		fileName, err := cpi.GetwxCodeUnlimit(scene, page)
		if err == nil {
			http.ServeFile(c.Response().Writer, c.Request(), fileName)
		} else {
			http.ServeFile(c.Response().Writer, c.Request(), fileName)
		}
		var err2 error
		return err2
	})

	// 获取推荐码(图片资源)
	e.GET("/qrcodebatch", func(c echo.Context) error {

		// 从多少开始
		offset, _ := strconv.Atoi(c.QueryParam("offset"))
		// 取生成多少个
		limit, _ := strconv.Atoi(c.QueryParam("limit"))

		for i := offset; i < limit+offset; i++ {
			// fmt.Println(i)
			cpi.GetwxCodeUnlimit(strconv.Itoa(i), `pages/index`)
		}

		return c.String(http.StatusOK, fmt.Sprint("qrcode batch init offset ", strconv.Itoa(offset), " limit ", strconv.Itoa(limit)))

	})

	e.GET("/gettodaysignusers", c.GetTodaySignUsers)

	// Restricted group
	api := e.Group("/api")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &c.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}

	api.Use(middleware.JWTWithConfig(config))
	// r.Use(middleware.JWT([]byte("secret")))

	// 获取用户信息
	api.GET("/user", c.GetUserInfo)
	api.GET("/checkopenid", c.CheckOpenID)
	api.GET("/getuserinfo", c.GetUserInfo)
	api.GET("/user/:id", c.GetUserSignInfo)
	api.GET("/user/:id/follow", c.GetTodaySignUsers)

	// 解密数据内容(保存用户数据到库)
	api.POST("/crypt", c.Crypt)

	// 签到
	api.GET("/dosign", c.UserDoSign)

	// 检查签到
	api.GET("/checksign", c.CheckUserSign)
	// 今天签到详细情况
	api.GET("/gettodaysigninfo", c.GetTodaySignInfo)

	api.GET("/getposter", c.GetPosterConfig)
	api.GET("/getappconfig", c.GetAppConfig)

	api.GET("/gettodaysignusers", c.GetTodaySignUsers)

	// 获取今日签到名单
	api.GET("/today", c.GetTodaySignUsers)

	// 获得用户排行月榜
	api.GET("/rank/month", c.GetMonthRank)

	// 获得用户排行周榜
	api.GET("/rank/week", c.GetWeekRank)

	// 获得用户排行总榜
	api.GET("/rank/all", c.GetAllRank)

	// 获取用户资源
	api.GET("/crypt", c.Crypt)
	// 图标
	e.File("favicon.ico", "images/favicon.ico")
	e.File("bg.jpg", "images/bg.jpg")

	e.Static("/static", "static")
	// e.Logger.Fatal(e.Start(":80"))
	// e.Logger.Fatal(e.Start(":8009"))
	e.Logger.Fatal(e.StartTLS(":443", "ssl/1781098_signapi.readfollow.com.pem", "ssl/1781098_signapi.readfollow.com.key"))
	// e.Logger.Fatal(e.StartAutoTLS(":443"))

}
