package models

import "time"

type Brg_kat struct {
	Kd_kat    int       `gorm:"type:int(5);primaryKey;autoIncrement"`
	Nm_kat    string    `gorm:"type:varchar(100);not null"`
	Aktif     string    `gorm:"type:enum('Y','T');default:Y"`
	On_create time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP()"`
	On_update time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP"`
}

type Tabler interface {
	TableName() string
}

func (Brg_kat) TableName() string {
	return "brg_kat"
}
