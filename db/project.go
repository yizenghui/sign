package db

import (
	"time"
)

//Project 签到项目
type Project struct {
	ID   uint   // `gorm:"primary_key"`
	Name string //`gorm:"index:xid"`     // 月键
	// Signs     []Sign
	// Users     []Fans // 粉丝与签到项目关联
	Intro     string
	Private   bool // 私有项目(检查粉丝表关联)
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// GetProjectsByIDS 获取今天签到的粉丝信息
func (p *Project) GetProjectsByIDS(ids []string) []Project {
	var list []Project
	DB().Where("id in (?)", ids).Limit(5).Offset(0).Order(`sign_at desc`).Find(&list)
	return list
}
