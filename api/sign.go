package api

import (
	"net/http"

	"github.com/labstack/echo"
	cpi "github.com/yizenghui/sign/core"
)

// NewSign 新建一个push
func NewSign(c echo.Context) error {
	openID := getOpenID(c)
	if openID != "" {
		state := c.QueryParam("state")
		f, _ := getUser(openID)
		if state == `complete` { // 完成的

		} // else fail

		return c.JSON(http.StatusOK, f)
	}
	return c.JSON(http.StatusOK, openID)
}

// NewPush 新建一个push
func NewPush(c echo.Context) error {
	openID := getOpenID(c)
	if openID != "" {
		f, _ := getUser(openID)
		return c.JSON(http.StatusOK, f)
	}
	return c.JSON(http.StatusOK, openID)
}

// GetPush 粉丝获取自己某个任务朋友push状态
func GetPush(c echo.Context) error {
	sessionKey := c.QueryParam("sk")
	encryptedData := c.QueryParam("ed")
	iv := c.QueryParam("iv")
	ret, _ := cpi.GetCryptData(sessionKey, encryptedData, iv)
	return c.JSON(http.StatusOK, ret)
}

// CheckPush 检查能不能push
func CheckPush(c echo.Context) error {
	sessionKey := c.QueryParam("sk")
	encryptedData := c.QueryParam("ed")
	iv := c.QueryParam("iv")
	ret, _ := cpi.GetCryptData(sessionKey, encryptedData, iv)
	return c.JSON(http.StatusOK, ret)
}
