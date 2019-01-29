package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo"
	cpi "github.com/yizenghui/sign/core"
	"github.com/yizenghui/sign/db"
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

// Poster 海报json数据
type Poster struct {
	Width  int64                    `json:"width"`
	Height int64                    `json:"height"`
	Clear  bool                     `json:"clear"`
	Views  []map[string]interface{} `json:"views"`
}

// BuildPoster 创建用户专属海报
func BuildPosterx(user *db.Fans) (Poster, error) {

	var MateArr = map[string]string{
		//
		"01": "一月",
		"02": "二月",
		"03": "三月",
		"04": "四月",
		"05": "五月",
		"06": "六月",
		"07": "七月",
		"08": "八月",
		"09": "九月",
		"10": "十月",
		"11": "十一月",
		"12": "十二月",
	}

	t := time.Now()
	// 背景
	mbg := map[string]interface{}{"type": "image", "url": "https://signapi.readfollow.com/static/images/bg.jpg", "top": 0, "left": 0, "width": 375, "height": 360}

	// 头像
	avatar := user.AvatarURL
	if avatar == `` {
		avatar = "https://wx.qlogo.cn/mmopen/vi_32/DYAIOgq83epJEPdPqQVgv6D8bojGT4DrGXuEC4Oe0GXs5sMsN4GGpCegTUsBgL9SPJkN9UqC1s0iakjQpwd4h4A/132"
	}
	muh := map[string]interface{}{"type": "image", "url": avatar, "top": 127.5, "left": 29, "width": 55, "height": 55}

	// 头像圈圈罩
	muhc := map[string]interface{}{"type": "image", "url": "https://signapi.readfollow.com/static/images/1531401349117.jpeg", "top": 127.5, "left": 29, "width": 55, "height": 55}

	// mo := t.Format("一月")
	d := t.Format(`01`)
	// 顶部月
	mmt := map[string]interface{}{"type": "text", "content": MateArr[d], "fontSize": 14, "color": "#402D16", "textAlign": "left", "top": 0, "left": 320, "bolder": false}

	// 顶部日
	mdt := map[string]interface{}{"type": "text", "content": t.Format(`02`), "fontSize": 22, "color": "#402D16", "textAlign": "left", "top": 18, "left": 320, "bolder": true}

	// 时间
	mts := map[string]interface{}{"type": "text", "content": t.Format(`15:04`), "fontSize": 26, "color": "#402D16", "textAlign": "left", "top": 128, "left": 96, "bolder": true}

	// 正文
	mct := map[string]interface{}{"type": "text", "content": "坚持自律", "fontSize": 16, "color": "#402D16", "textAlign": "left", "top": 138, "left": 176, "bolder": false}

	// 坚持天数
	mcd := map[string]interface{}{"type": "text", "content": user.AllToT, "fontSize": 26, "color": "red", "textAlign": "center", "top": 128, "left": 270, "bolder": true}

	// 正文
	mctt := map[string]interface{}{"type": "text", "content": "天", "fontSize": 16, "color": "#402D16", "textAlign": "left", "top": 138, "left": 302, "bolder": false}

	// 正文
	mtt := map[string]interface{}{"type": "text", "content": "124人正在参与", "fontSize": 16, "color": "#383549", "textAlign": "left", "top": 168, "left": 96, "bolder": false}

	// 用户分享二维码
	muqr := map[string]interface{}{"type": "image", "url": "https://signapi.readfollow.com/static/images/1531385433625.jpeg", "top": 250, "left": 265, "width": 68, "height": 68}

	// 正文
	mqrt := map[string]interface{}{"type": "text", "content": "扫码一起改变", "fontSize": 14, "color": "#383549", "textAlign": "left", "top": 320, "left": 255, "bolder": false}

	poster := Poster{
		375,
		360,
		true,
		[]map[string]interface{}{
			mbg,
			muh,
			muhc,
			mmt,
			mdt,
			mts,
			mct,
			mcd,
			mctt,
			mtt,
			muqr,
			mqrt,
		},
	}
	return poster, nil
}

// BuildPoster 创建用户专属海报
func BuildPoster(user *db.Fans) (Poster, error) {

	var MateArr = map[string]string{
		//
		"01": "一月",
		"02": "二月",
		"03": "三月",
		"04": "四月",
		"05": "五月",
		"06": "六月",
		"07": "七月",
		"08": "八月",
		"09": "九月",
		"10": "十月",
		"11": "十一月",
		"12": "十二月",
	}

	// 头像
	avatar := user.AvatarURL
	if avatar == `` { // todo 给默认头像
		avatar = "https://wx.qlogo.cn/mmopen/vi_32/DYAIOgq83epJEPdPqQVgv6D8bojGT4DrGXuEC4Oe0GXs5sMsN4GGpCegTUsBgL9SPJkN9UqC1s0iakjQpwd4h4A/132"
	}

	avatar = strings.Replace(avatar, `/132`, `/0`, -1)

	muh := map[string]interface{}{"type": "image", "url": avatar, "top": 410, "left": 84, "width": 200, "height": 200}

	t := time.Now()
	// 背景https://signapi.readfollow.com/static/images/20190128164201.jpg
	mbg := map[string]interface{}{"type": "image", "url": "https://i.loli.net/2019/01/29/5c4ff12a37b95.jpg", "top": 0, "left": 0, "width": 1080, "height": 1080}

	// 遮罩层 https://signapi.readfollow.com/static/images/mask.png
	mmg := map[string]interface{}{"type": "image", "url": "https://i.loli.net/2019/01/29/5c4ff129e5088.png", "top": 0, "left": 0, "width": 1080, "height": 1080}

	// mo := t.Format("一月")
	d := t.Format(`01`)
	// 顶部月
	mmt := map[string]interface{}{"type": "text", "content": MateArr[d], "fontSize": 42, "color": "#ffffff", "textAlign": "left", "top": 0, "left": 910, "bolder": false}

	// 顶部日
	mdt := map[string]interface{}{"type": "text", "content": t.Format(`02`), "fontSize": 72, "color": "#ffffff", "textAlign": "left", "top": 42, "left": 910, "bolder": true}

	// 时间
	mts := map[string]interface{}{"type": "text", "content": t.Format(`15:04`), "fontSize": 80, "color": "#402D16", "textAlign": "left", "top": 422, "left": 330, "bolder": true}

	// 正文
	mct := map[string]interface{}{"type": "text", "content": "坚持自律", "fontSize": 50, "color": "#402D16", "textAlign": "left", "top": 450, "left": 590, "bolder": false}

	// 坚持天数
	mcd := map[string]interface{}{"type": "text", "content": user.AllToT, "fontSize": 80, "color": "red", "textAlign": "center", "top": 422, "left": 830, "bolder": true}

	// 正文
	mctt := map[string]interface{}{"type": "text", "content": "天", "fontSize": 50, "color": "#402D16", "textAlign": "left", "top": 450, "left": 870, "bolder": false}

	// 正文
	mtt := map[string]interface{}{"type": "text", "content": fmt.Sprint(user.AllFansCount(), "人正在参与"), "fontSize": 40, "color": "#383549", "textAlign": "left", "top": 560, "left": 340, "bolder": false}

	//
	// 用户分享二维码
	qrfile, err := cpi.GetwxCodeUnlimit(strconv.FormatInt(int64(user.ID), 10), `pages/index`)

	muqr := map[string]interface{}{"type": "image", "url": fmt.Sprint("https://signapi.readfollow.com/", qrfile), "top": 920, "left": 900, "width": 150, "height": 150}
	if err != nil {

		// 如果生成带参数二维码出错,设置回默认的二维码
		muqr = map[string]interface{}{"type": "image", "url": "https://signapi.readfollow.com/static/images/qrcode.jpg", "top": 920, "left": 900, "width": 150, "height": 150}

	}
	// 鸡汤
	mhc := map[string]interface{}{"type": "text", "content": "真正的自由，是自我掌握而不是随心所欲", "fontSize": 32, "color": "#827f7b", "textAlign": "right", "top": 950, "left": 860, "bolder": false}

	// 正文
	mqrt := map[string]interface{}{"type": "text", "content": "扫码一起自律", "fontSize": 32, "color": "#827f7b", "textAlign": "right", "top": 1000, "left": 860, "bolder": false}

	poster := Poster{
		1080,
		1080,
		true,
		[]map[string]interface{}{
			mbg,
			muh,
			mmg,
			mmt,
			mdt,
			mts,
			mct,
			mcd,
			mctt,
			mtt,
			muqr,
			mqrt,
			mhc,
		},
	}
	return poster, nil
}

// GetPosterConfig 获取生成海报的数据配置
func GetPosterConfig(c echo.Context) error {
	openID := getOpenID(c)
	if openID != "" {
		user, _ := getUser(openID)
		poster, _ := BuildPoster(user)
		return c.JSON(http.StatusOK, poster)
	}
	return c.JSON(http.StatusUnauthorized, `openid is empty.`)
}

// GetAppConfig 获取生成海报的数据配置
func GetAppConfig(c echo.Context) error {
	openID := getOpenID(c)
	if openID != "" {
		user, _ := getUser(openID)

		config := map[string]interface{}{
			"title":         "每天自律打卡",
			"sharetitle":    "自律改变人生",
			"sharepath":     fmt.Sprint("/pages/index?from=", user.ID),
			"shareimageUrl": "",
			// "shareimageUrl": "https://signapi.readfollow.com/static/images/bg.jpg",
		}

		return c.JSON(http.StatusOK, config)
	}
	return c.JSON(http.StatusUnauthorized, `openid is empty.`)
}

// GetTodaySignUsers 获取今日签到用户信息
func GetTodaySignUsers(c echo.Context) error {
	// openID := getOpenID(c)
	page, _ := strconv.Atoi(c.QueryParam("page"))

	data := cpi.GetTodaySignUsers(int64(page))
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
