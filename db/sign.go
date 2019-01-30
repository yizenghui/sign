package db

import (
	"time"
)

//Sign 签到记录
type Sign struct {
	ID     uint  // `gorm:"primary_key"`
	FansID uint  //`gorm:"index:fans_id"` // 粉丝id
	MID    int64 //`gorm:"index:xid"`     // 月键
	WID    int64 //`gorm:"index:xid"`     // 周键
	DID    int64 //`gorm:"index:xid"`     // 日键
	Score  int64
	PIDS   string
	// Projects  []Project
	CreatedAt time.Time
}
