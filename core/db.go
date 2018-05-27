package core

import (
	// "fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Fans 粉丝
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

//Share 分享记录
type Share struct {
	ID        uint `gorm:"primary_key"`
	FansID    uint `gorm:"index:index"`
	TaskID    uint `gorm:"index:index"`
	CreatedAt time.Time
}

// Task 提交的url
type Task struct {
	ID               uint      `gorm:"primary_key"`
	FansID           uint      //发起人id
	City             string    `sql:"index"`                // 城市(发起人所在的)
	Title            string    `gorm:"type:varchar(64);"`   // 活动标题
	Intro            string    `gorm:"type:varchar(1024);"` // 活动描述
	Statement        string    `gorm:"type:varchar(1024);"` // 声明
	TotalNum         int64     //总访问量
	Number           int64     //最大可获奖人数
	CompletionNumber int64     //当前完成人数
	Fore             int64     //最低获取火力条件
	StartAt          time.Time //开始时间
	EndAt            time.Time //结束时间
	Images           string    `gorm:"type:text;"` // 展品图片
	CreatedAt        time.Time
	UpdatedAt        time.Time
	SpreadAt         *time.Time `sql:"index:date"` //推广期截止时间
	ModeratedAt      *time.Time `sql:"index:date"` //审核时间
	DeletedAt        *time.Time `sql:"index:date"`
}

//Confirm 粉丝参加活动结账后帮其证实真实有效
type Confirm struct {
	ID        uint       `gorm:"primary_key"`
	FansID    uint       `gorm:"index:id"` //粉丝ID
	TaskID    uint       `gorm:"index:id"` //活动ID
	DeletedAt *time.Time `sql:"index"`
}

//Report 粉丝举报活动
type Report struct { // 举报需知： 受理存在以下情况的活动，虚假、挂羊头卖狗肉、额外收费或条件直接影响活动结算的
	ID        uint   `gorm:"primary_key"`
	FansID    uint   `gorm:"index:id"`            //粉丝ID
	TaskID    uint   `gorm:"index:id"`            //活动ID
	Intro     string `gorm:"type:varchar(1024);"` //描述
	Images    string `gorm:"type:text;"`          //证图
	State     int16  //待处理 已撤消 受理中 协商解决 实锤 假锤
	Reply     string `gorm:"type:varchar(1024);"` //答复
	Contact   string //留下联系方式 微信号或者手机号
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

//Join 粉丝参加活动成绩
type Join struct {
	ID           uint      `gorm:"primary_key"`
	Fore         uint      `gorm:"index:id"` //有效火力
	FansID       uint      `gorm:"index:id"` //粉丝ID
	TaskID       uint      `gorm:"index:id"` //活动ID
	ReachAt      time.Time //达成时间
	SettlementAt time.Time //结算时间
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `sql:"index"`
}

//Push 助力记录
type Push struct {
	ID        uint `gorm:"primary_key"`
	FansID    uint `gorm:"index:user_id"` // 谁在收集助力
	JoinID    uint `gorm:"index:user_id"` // 参加活动凭证
	PushID    uint `gorm:"index:user_id"` // 哪个朋友来给助力
	CreatedAt time.Time
}

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

// Post 提交的url
type Post struct {
	ID               uint   `gorm:"primary_key"`
	Title            string `gorm:"type:varchar(1024);"`             // 微信文章地址
	URL              string `gorm:"type:varchar(1024);unique_index"` // 微信文章地址
	SubNum           int64  // 订阅人次 用户每提交一次+1
	FolNum           int64  // 当前关注人数 注，如果有人关注，每过8小时检查更新
	ShareNum         int64  `gorm:"index"`      // 分享次数 现在用于排序
	ChapterFragments string `gorm:"type:text;"` // 章节片段
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time `sql:"index"`
}

// Subscribe 粉丝
type Subscribe struct {
	ID        uint   `gorm:"primary_key"`
	FansID    uint   `sql:"index"`               //粉丝 ID
	OpenID    string `gorm:"type:varchar(255);"` //提交的openid
	PostID    uint   `sql:"index"`               //post ID
	FormID    string `gorm:"type:varchar(255);"` //订阅formID，一次订阅只能推送一次通知
	Push      bool   //是否推送
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"` //删除后不再推送
}

var db *gorm.DB

//DB 返回 *gorm.DB
func DB() *gorm.DB {
	if db == nil {

		newDb, err := newDB()
		if err != nil {
			panic(err)
		}
		newDb.DB().SetMaxIdleConns(10)
		newDb.DB().SetMaxOpenConns(100)

		newDb.LogMode(false)
		db = newDb
	}

	return db
}

func newDB() (*gorm.DB, error) {

	// sqlConnection := fmt.Sprintf(
	// 	"host=%v user=%v port=%v dbname=%v sslmode=%v password=%v",
	// 	config.Database.Host,
	// 	config.Database.User,
	// 	config.Database.Port,
	// 	config.Database.Dbname,
	// 	config.Database.Sslmode,
	// 	config.Database.Password,
	// )
	// db, err := gorm.Open(config.Database.Type, sqlConnection)
	db, err := gorm.Open("sqlite3", "fireapi.db")

	if err != nil {
		return nil, err
	}
	return db, nil
}
