package models

import (
	"apps_barang/config"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Kd_user   int       `gorm:"type:int(5);primaryKey;autoIncrement" json:"kd_user"`
	Nm_user   string    `gorm:"type:varchar(100);not null" json:"nm_user"`
	Username  string    `gorm:"type:varchar(100);not null" json:"username"`
	Salt      string    `gorm:"type:text;not null" json:"salt"`
	Password  string    `gorm:"type:text;not null" json:"password"`
	On_create time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP()" json:"on_create"`
	On_update time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP" json:"on_update"`
}

func (User) TableName() string {
	return "user"
}
func FindUser(field []string, where map[string]interface{}) *gorm.DB {
	Mysql := config.InitMysql()
	Build := Mysql.Model(&User{})
	Build.Select(field)
	//Build = JoinBrg(Build)
	for field, value := range where {
		if value == false {
			Build.Where(field)
		} else {
			Build.Where(field, value)
		}
	}
	return Build
}
func CountUser(where map[string]interface{}) int64 {
	var Counta int64
	Build := FindUser([]string{
		"COUNT(user.kd_user) AS counta",
	}, where)
	Build.Scan(&Counta)
	return Counta
}
