package entity

import (
	"time"
)

type FamilyList struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"fl_id"`
	CustomerID uint      `gorm:"not null;index" json:"cst_id"`
	Name       string    `gorm:"type:varchar(100);not null" json:"fl_name"`
	DOB        time.Time `gorm:"type:date;not null" json:"fl_dob"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Customer   Customer  `gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
}

func (FamilyList) TableName() string { return "family_list" }
