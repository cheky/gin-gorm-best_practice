package models

import "time"

type Brg_kat struct {
	Kd_kat    int       `gorm:"type:int(5);primaryKey;autoIncrement" json:"kd_kat"`
	Nm_kat    string    `gorm:"type:varchar(100);not null" json:"nm_kat"`
	Aktif     string    `gorm:"type:enum('Y','T');default:Y" json:"aktif"`
	On_create time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP()" json:"on_create"`
	On_update time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP" json:"on_update"`
}

type Tabler interface {
	TableName() string
}

func (Brg_kat) TableName() string {
	return "brg_kat"
}
