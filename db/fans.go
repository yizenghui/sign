package db

import (
	"time"
)

// Fans 粉丝数据信息
type Fans struct {
	ID         uint   `gorm:"primary_key"`
	OpenID     string `gorm:"type:varchar(255);unique_index"`
	UnionID    string
	NickName   string
	Gender     int
	City       string
	Province   string
	Country    string
	AvatarURL  string
	Language   string
	Timestamp  int64
	Trust      int16 `gorm:"default:30;"` // 信任度
	AppID      string
	SessionKey string // 粉丝上次的session key 如果有变化，同步一次粉丝数据
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
}

// GetFansByOpenID  通过openID获取粉丝信息如果没有的话进行初始化
func (fans *Fans) GetFansByOpenID(openID string) {
	DB().Where(Fans{OpenID: openID}).FirstOrCreate(&fans)
}

// Save 保存粉丝信息
func (fans *Fans) Save() {
	DB().Save(&fans)
}

// ShareLog 记录 openID 分享 Task 记录
func (fans *Fans) ShareLog() {
	var share = Share{}
	share.FansID = fans.ID
	DB().Create(share)
}
