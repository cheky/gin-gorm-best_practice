package models

import (
	"apps_barang/config"
	"time"

	"gorm.io/gorm"
)

type Brg struct {
	Kd_brg    int       `gorm:"type:bigint(20);primaryKey;autoIncrement" json:"kd_brg"`
	Nm_brg    string    `gorm:"type:tinytext;not null" json:"nm_brg"`
	Kd_kat    int       `gorm:"type:int(5);null" json:"kd_kat"`
	Aktif     string    `gorm:"type:enum('Y','T');default:Y" json:"aktif"`
	On_create time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP()" json:"on_create"`
	On_update time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP" json:"on_update"`
}

func (Brg) TableName() string {
	return "brg"
}
func JoinBrg(db *gorm.DB) *gorm.DB {
	db.Joins("INNER JOIN brg_kat ON brg_kat.kd_kat = brg.kd_kat")
	return db
}
func InsertBrg(brg *Brg) (*gorm.DB, int) {
	Mysql := config.InitMysql()
	Result := Mysql.Create(&brg)
	return Result, brg.Kd_brg
}
func DeleteBrg(where map[string]interface{}) *gorm.DB {
	Mysql := config.InitMysql()
	Build := Mysql.Model(&Brg{})
	for field, value := range where {
		if value == false {
			Build.Where(field)
		} else {
			Build.Where(field, value)
		}
	}
	Build.Delete(&Brg{})
	return Build
}
func UpdateBrg(brg *Brg, where map[string]interface{}) *gorm.DB {
	Mysql := config.InitMysql()
	Build := Mysql.Model(&Brg{})
	for field, value := range where {
		if value == false {
			Build.Where(field)
		} else {
			Build.Where(field, value)
		}
	}
	Build.Updates(brg)
	return Build
}

func FindBrg(field []string, where map[string]interface{}) *gorm.DB {
	Mysql := config.InitMysql()
	Build := Mysql.Model(&Brg{})
	Build.Select(field)
	Build = JoinBrg(Build)
	for field, value := range where {
		if value == false {
			Build.Where(field)
		} else {
			Build.Where(field, value)
		}
	}
	return Build
}
func CountBrg(where map[string]interface{}) int64 {
	var Counta int64
	Build := FindBrg([]string{
		"COUNT(brg.kd_brg) AS counta",
	}, where)
	Build.Scan(&Counta)
	return Counta
}
