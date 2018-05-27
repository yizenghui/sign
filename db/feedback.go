package db

import (
	"time"
)

// Feedback 粉丝
type Feedback struct {
	ID        uint   `gorm:"primary_key"`
	OpenID    string `gorm:"type:varchar(255);index"` // 微信文章地址
	FormID    string `gorm:"type:varchar(255);"`      //订阅formID，一次订阅只能推送一次通知
	Problem   string `gorm:"type:text;"`              // 问题
	Answer    string `gorm:"type:text;"`              // 答复
	Show      bool   //是否显示
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// Save Feedback
func (feedback *Feedback) Save() {
	DB().Save(&feedback)
}
