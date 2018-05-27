package api

// import (
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"github.com/labstack/echo"
// 	cpi "github.com/yizenghui/fire/core"
// 	"github.com/yizenghui/sign/db"
// )

// // NewTask 创建新任务 post
// // /newstask post
// func NewTask(c echo.Context) error {

// 	fans, e := getUser(getOpenID(c))
// 	if e != nil {
// 		return echo.ErrUnauthorized
// 	}

// 	title := c.FormValue("title")
// 	intro := c.FormValue("intro")
// 	if title != `` {

// 		t := db.Task{
// 			FansID: fans.ID,
// 			City:   fans.City,
// 			Title:  title,
// 			Intro:  intro,
// 		}
// 		db.DB().Create(&t)

// 		return c.JSON(http.StatusOK, t)
// 	}
// 	return echo.ErrUnauthorized
// }

// // GetTaskInfo 获取一个活动详细
// // taskinfo/:id  get
// func GetTaskInfo(c echo.Context) error {
// 	// fans, e := getUser(getOpenID(c))
// 	// if e != nil {
// 	// 	return echo.ErrUnauthorized
// 	// }
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	// id, _ := strconv.Atoi(c.QueryParam("id"))
// 	var t = cpi.Task{}
// 	t.GetTaskByID(int64(id))

// 	return c.JSON(http.StatusOK, t)
// }

// // JoinTask 加入一个活动
// // jointask/:id  get
// func JoinTask(c echo.Context) error {
// 	fans, e := getUser(getOpenID(c)) // 获取fans信息
// 	if e != nil {
// 		return echo.ErrUnauthorized
// 	}
// 	id, _ := strconv.Atoi(c.Param("id")) //获取task id
// 	var t = cpi.Task{}
// 	t.GetTaskByID(int64(id)) // 获取task t

// 	nowdate := time.Now() // 当前时间

// 	if t.ID > 0 && nowdate.After(t.StartAt) && nowdate.Before(t.EndAt) { // 有id 后于开始时间 前于结束时间 (活动期限内)
// 		// if t.StartAt
// 		var join = cpi.Join{}
// 		cpi.DB().Where(&cpi.Join{TaskID: uint(id), FansID: fans.ID}).First(join)
// 		if join.ID == 0 {
// 			// 新增
// 			newJoin := cpi.Join{
// 				TaskID: uint(id),
// 				FansID: fans.ID,
// 			}
// 			cpi.DB().Create(&newJoin)
// 			return c.JSON(http.StatusOK, newJoin)
// 		}
// 	}
// 	return c.JSON(http.StatusOK, t)
// }

// // CheckJoin 检查能否加入一个活动
// // checkjoin/:id  get
// func CheckJoin(c echo.Context) error {
// 	fans, e := getUser(getOpenID(c)) // 获取fans信息
// 	if e != nil {
// 		return echo.ErrUnauthorized
// 	}
// 	id, _ := strconv.Atoi(c.Param("id")) //获取task id
// 	var t = cpi.Task{}
// 	t.GetTaskByID(int64(id)) // 获取task t

// 	nowdate := time.Now() // 当前时间

// 	if t.ID > 0 && nowdate.After(t.StartAt) && nowdate.Before(t.EndAt) { // 有id 后于开始时间 前于结束时间 (活动期限内)
// 		// if t.StartAt
// 		var join = cpi.Join{}
// 		cpi.DB().Where(&cpi.Join{TaskID: uint(id), FansID: fans.ID}).First(join)
// 		if join.ID == 0 {
// 			// 新增
// 			newJoin := cpi.Join{
// 				TaskID: uint(id),
// 				FansID: fans.ID,
// 			}
// 			cpi.DB().Create(&newJoin)
// 			return c.JSON(http.StatusOK, newJoin)
// 		}
// 		return c.JSON(http.StatusOK, join)
// 	}
// 	return c.JSON(http.StatusOK, t)
// }
