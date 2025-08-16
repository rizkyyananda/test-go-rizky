package entity

import (
	"time"
)

type Nationality struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"nationality_id"`
	Name      string     `gorm:"type:varchar(100);not null" json:"nationality_name"`
	Code      string     `gorm:"type:varchar(10);not null;uniqueIndex" json:"nationality_code"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	Customers []Customer `gorm:"foreignKey:NationalityID" json:"customers,omitempty"`
}

func (Nationality) TableName() string { return "nationality" }
