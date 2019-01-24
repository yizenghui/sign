package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	cpi "github.com/yizenghui/sign/core"
)

// CheckTodaySign 今天能否签到
func CheckTodaySign(c echo.Context) error {
	openID := getOpenID(c)
	if openID != "" {
		if cpi.CheckOpenIDCanSign(openID) == nil {
			return c.JSON(http.StatusOK, `t`)
		}
		return c.JSON(http.StatusOK, `f`)
	}
	return c.JSON(http.StatusOK, `f`)
}

// DoSign 今天能否签到
func DoSign(c echo.Context) error {
	openID := getOpenID(c)
	if openID != "" {
		if cpi.FansDoSign(openID) == nil {
			return c.JSON(http.StatusOK, `t`)
		}
		return c.JSON(http.StatusOK, `f`)
	}
	return c.JSON(http.StatusOK, `f`)
}

// GetTodayNewSign 获取今日最新签到
func GetTodayNewSign(c echo.Context) error {
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

// GetSignRanking 用户签到排名(总榜)
func GetSignRanking(c echo.Context) error {
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

// GetSignLasting 最大持续天数
func GetSignLasting(c echo.Context) error {
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

// GetUserInfo 新建一个push
func GetUserInfo(c echo.Context) error {
	openID := getOpenID(c)
	if openID != "" {
		f, _ := getUser(openID)
		return c.JSON(http.StatusOK, f)
	}
	return c.JSON(http.StatusOK, openID)
}

// UserDoSign 用户进行签到
func UserDoSign(c echo.Context) error {
	openID := getOpenID(c)
	if openID != "" {
		if cpi.FansDoSign(openID) == nil {
			return c.JSON(http.StatusOK, `t`)
		}
		return c.JSON(http.StatusNotFound, `Please come back tomorrow.`)
	}
	return c.JSON(http.StatusUnauthorized, `openid is empty.`)
}

// GetUserSignInfo 查看用户详细 用户详情页
func GetUserSignInfo(c echo.Context) error {
	openID := getOpenID(c)
	if openID != "" {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, cpi.GetFansInfo(uint(id)))
	}
	return c.JSON(http.StatusUnauthorized, `openid is empty.`)
}

// CheckUserSign 检查用户今日签到情况
func CheckUserSign(c echo.Context) error {
	openID := getOpenID(c)
	if openID != "" {
		if cpi.CheckOpenIDCanSign(openID) == nil {
			return c.JSON(http.StatusOK, `t`)
		}
		return c.JSON(http.StatusOK, `Please come back tomorrow.`)
	}
	return c.JSON(http.StatusUnauthorized, `openid is empty.`)
}

// GetTodaySignInfo 获取今天签到详细情况
func GetTodaySignInfo(c echo.Context) error {
	openID := getOpenID(c)
	if openID != "" {
		return c.JSON(http.StatusOK, cpi.GetTodaySignInfo(openID))
	}
	return c.JSON(http.StatusUnauthorized, `openid is empty.`)
}

// GetTodaySignUsers 获取今日签到用户信息
func GetTodaySignUsers(c echo.Context) error {
	// openID := getOpenID(c)
	data := cpi.GetTodaySignUsers(`openID`)
	return c.JSON(http.StatusOK, data)
}

// GetMonthRank 获取月排行榜
func GetMonthRank(c echo.Context) error {
	openID := getOpenID(c)
	if openID != "" {
		f, _ := getUser(openID)
		return c.JSON(http.StatusOK, f)
	}
	return c.JSON(http.StatusOK, openID)
}

// GetWeekRank 获取周排行榜
func GetWeekRank(c echo.Context) error {
	openID := getOpenID(c)
	if openID != "" {
		f, _ := getUser(openID)
		return c.JSON(http.StatusOK, f)
	}
	return c.JSON(http.StatusOK, openID)
}

// GetAllRank 获取总排行榜
func GetAllRank(c echo.Context) error {
	openID := getOpenID(c)
	if openID != "" {
		f, _ := getUser(openID)
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
