package db

import (
	"fmt"
	"strconv"
	"time"
)

// Fans 粉丝数据信息
type Fans struct {
	ID          uint   `gorm:"primary_key"`
	OpenID      string `gorm:"type:varchar(255);unique_index"`
	UnionID     string
	NickName    string
	Gender      int
	City        string
	Province    string
	Country     string
	AvatarURL   string
	Language    string
	Timestamp   int64 // 微信用户数据上的时间戳
	AllToT      int64 // 总统计
	MonthToT    int64 // 月签统计
	WeekToT     int64 // 周签统计
	MonthSignID int64 // 月签ID
	WeekSignID  int64 // 周签ID
	DaySignID   int64 // 上次签得时间 (用于判断是否断签)
	AllRank     int64 // 总排行分
	MonthRank   int64 // 月排行分
	WeekRank    int64 // 周排行分
	LastSignAt  int64 // 下次签到有加成日期Ymd (当前日期+1判断连签)
	SignScore   int16 // 签到得分(额外加成)
	NextSignAdd int16 // 下次签到加成 分享被浏览+1 (签到时清零)
	PubAt       int64 // 设置公开我的签到情况时间(不为空时公开签到用户信息)
	AppID       string
	SessionKey  string    // 粉丝上次的session key 如果有变化，同步一次粉丝数据
	SignAt      time.Time `sql:"index"` //最后一次签到时间
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}

// GetFansByOpenID  通过openID获取粉丝信息如果没有的话进行初始化
func (fans *Fans) GetFansByOpenID(openID string) {
	DB().Where(Fans{OpenID: openID}).FirstOrCreate(&fans)
}

// GetFansByID  通过ID获取粉丝信息
func (fans *Fans) GetFansByID(id uint) {
	DB().Where(Fans{ID: id}).First(&fans)
}

// Save 保存粉丝信息
func (fans *Fans) Save() {
	DB().Save(&fans)
}

// AllFans 获取粉丝所有信息(测试用)
func (fans *Fans) AllFans() []Fans {
	var list []Fans
	DB().Find(&list)
	return list
}

// AllFansCount 获取粉丝所有信息(测试用)
func (fans *Fans) AllFansCount() (count int64) {
	DB().Model(&Fans{}).Count(&count)
	return count
}

// GetTodaySignFans 获取今天签到的粉丝信息
func (fans *Fans) GetTodaySignFans() []Fans {
	var list []Fans
	DB().Limit(10).Offset(0).Order(`sign_at desc`).Find(&list)
	return list
}

// GetLastSign 获取粉丝最后一次签到信息
func (fans *Fans) GetLastSign() Sign {
	var sign Sign
	DB().Where("fans_id = ?", fans.ID).Order(`id desc`).First(&sign)
	return sign
}

// GetTodaySignFansPage 获取今天签到的粉丝信息
func (fans *Fans) GetTodaySignFansPage(page int64) []Fans {
	_, _, did := XID(time.Now())
	offset := (page - 1) * 20
	if offset < 0 {
		offset = 0
	}
	var list []Fans
	DB().Where(`day_sign_id = ?`, did).Limit(20).Offset(offset).Order(`sign_at desc`).Find(&list)
	return list
}

// GetTodaySignFansW 获取今天签到的粉丝信息
func (fans *Fans) GetTodaySignFansW(signAt time.Time) []Fans {
	_, _, did := XID(time.Now())
	var list []Fans
	DB().Where(`sign_at < ?`, signAt).Where(`day_sign_id = ?`, did).Limit(10).Offset(0).Order(`sign_at desc`).Find(&list)
	return list
}

// ShareLog 记录 openID 分享 Task 记录
func (fans *Fans) ShareLog() {
	var share = Share{}
	share.FansID = fans.ID
	DB().Create(&share)
}

// Relation 记录关联 粉丝id
func (fans *Fans) Relation(relationID int) {
	DB().Where(Relation{FansID: fans.ID, RelationID: uint(relationID)}).FirstOrCreate(&fans)
}

// XID 分拆三个时间ID
func XID(then time.Time) (mid, wid, did int64) {
	var mstr, wstr, dstr string
	mstr = fmt.Sprintf(`100%v`, then.Format("200601"))
	dstr = fmt.Sprintf(`3%v`, then.Format("20060102"))
	y, w := then.ISOWeek()
	if len(string(w)) == 1 {
		wstr = fmt.Sprintf(`200%v%v`, y, w)
	} else {
		wstr = fmt.Sprintf(`200%v0%v`, y, w)
	}
	month, _ := strconv.Atoi(mstr) // mid 是由年+月组成的整数
	week, _ := strconv.Atoi(wstr)  // wid 是由年+第几周组成的整数
	day, _ := strconv.Atoi(dstr)   // wid 是由年+第几周组成的整数
	mid, wid, did = int64(month), int64(week), int64(day)
	return
}

// DoSign 进行签到
func (fans *Fans) DoSign(pids string) bool {
	if !fans.CheckSign() { // 首先检查一下今天能不能签到
		return false
	}
	// 获得签到得分
	score := fans.GetThenSignScore()
	// 粉丝现在是连签
	if fans.GetThenSignIsAddition() {
		if fans.SignScore < 99 {
			fans.SignScore++
		}
	} else {
		// 断签加成减半
		if fans.SignScore > 0 {
			fans.SignScore = fans.SignScore / 2
		}
		fans.SignScore++
	}
	mid, wid, did := XID(time.Now())
	if fans.MonthSignID == mid { // 是本月签到
		fans.MonthRank = fans.MonthRank + score
		fans.MonthToT++
	} else {
		fans.MonthSignID = mid
		fans.MonthRank = score
		fans.MonthToT = 1
	}

	if fans.WeekSignID == wid { // 是本周签到
		fans.WeekRank = fans.WeekRank + score
		fans.WeekToT++
	} else {
		fans.WeekSignID = wid
		fans.WeekRank = score
		fans.WeekToT = 1
	}
	fans.AllRank = fans.AllRank + score
	fans.AllToT++
	fans.DaySignID = did
	fans.SignAt = time.Now()             // 记录签到时间
	fans.NextSignAdd = 0                 // 重置下次加成
	fans.LastSignAt = fans.GetNextSign() // 记录下次连签时间
	fans.SignLog(score, pids)
	fans.Save()
	return true
}

// SignLog 保存签到记录
func (fans *Fans) SignLog(score int64, pids string) (sign Sign) {
	mid, wid, did := XID(time.Now())
	// 签到记录
	sign.FansID = fans.ID
	sign.MID = mid
	sign.WID = wid
	sign.DID = did
	sign.Score = score // 本次签到得分
	sign.PIDS = pids
	DB().Create(&sign)
	return sign
}

// GetThenSignScore 计算现在签到得分
func (fans *Fans) GetThenSignScore() int64 {
	if fans.GetThenSignIsAddition() {
		return int64(1 + fans.NextSignAdd)
	}
	return int64(1 + fans.SignScore + fans.NextSignAdd)
}

// GetThenSignIsAddition 今天是连签?
func (fans *Fans) GetThenSignIsAddition() bool {
	todayStr := time.Now().Format("20060102")
	return string(fans.LastSignAt) == todayStr
}

// GetNextSign 获取下次连签日期
func (fans *Fans) GetNextSign() int64 {
	tomorrowStr := time.Now().AddDate(0, 0, 1).Format("20060102")
	tomorrow, _ := strconv.Atoi(tomorrowStr)
	return int64(tomorrow)
}

// CheckSign 检查现在能否签到
func (fans *Fans) CheckSign() bool {
	_, _, did := XID(time.Now())
	if fans.DaySignID < did {
		return true
	}
	return false
}
