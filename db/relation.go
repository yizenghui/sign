package db

import (
	"time"
)

// Relation 关联关系(粉丝分享小程序吸引别人来)
type Relation struct {
	ID         uint `gorm:"primary_key"`
	FansID     uint `gorm:"index:id"` //粉丝ID
	RelationID uint `gorm:"index:id"` //被吸来ID
	CreatedAt  time.Time
}

// RelationLog 记录 openID 吸引粉丝访问
func (fans *Fans) RelationLog(relationID int) {
	var rela = Relation{}
	rela.FansID = fans.ID
	rela.RelationID = uint(relationID)
	DB().Create(rela)
}
