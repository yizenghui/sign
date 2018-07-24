package main

import (
	"net/http"

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

	// 获取用户签名
	e.GET("/gettoken", c.GetToken)
	// 解密数据内容(保存数据到库)
	e.GET("/crypt", c.Crypt)

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

	// Restricted group
	api := e.Group("/api")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &c.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	api.Use(middleware.JWTWithConfig(config))
	// r.Use(middleware.JWT([]byte("secret")))

	// 新增助力
	api.POST("/sign", c.NewPush)

	// 获取用户资源
	api.GET("/crypt", c.Crypt)
	// 图标
	e.File("favicon.ico", "images/favicon.ico")
	e.Logger.Fatal(e.Start(":8009"))
	// e.Logger.Fatal(e.StartAutoTLS(":443"))

}
