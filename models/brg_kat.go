package models

import (
	"apps_barang/config"
	"time"

	"gorm.io/gorm"
)

type Brg_kat struct {
	Kd_kat    int       `gorm:"type:int(5);primaryKey;autoIncrement" json:"kd_kat"`
	Nm_kat    string    `gorm:"type:varchar(100);not null" json:"nm_kat"`
	Aktif     string    `gorm:"type:enum('Y','T');default:Y" json:"aktif"`
	On_create time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP()" json:"on_create"`
	On_update time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP" json:"on_update"`
	Brg       []Brg     `gorm:"foreignKey:kd_kat;references:kd_kat;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

func (Brg_kat) TableName() string {
	return "brg_kat"
}
func InsertBrg_kat(brg_kat *Brg_kat) (*gorm.DB, *Brg_kat) {
	Mysql := config.InitMysql()
	Result := Mysql.Create(&brg_kat)
	return Result, brg_kat
}
func UpdateBrg_kat() *gorm.DB {
	Mysql := config.InitMysql()
	Result := Mysql.Model(Brg_kat{})
	return Result
}

func FindBrg_kat(field []string) *gorm.DB {
	Mysql := config.InitMysql()
	Build := Mysql.Model(&Brg_kat{})
	Build.Select(field)
	return Build
}
