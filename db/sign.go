package db

import (
	"time"
)

//Sign 签到记录
type Sign struct {
	ID        uint `gorm:"primary_key"`
	FansID    uint `gorm:"index:user_id"` // 谁在收集助力
	JoinID    uint `gorm:"index:user_id"` // 参加活动凭证
	PushID    uint `gorm:"index:user_id"` // 哪个朋友来给助力
	CreatedAt time.Time
}
