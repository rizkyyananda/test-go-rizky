package entity

import (
	"time"
)

type Customer struct {
	ID            uint         `gorm:"primaryKey;autoIncrement" json:"cst_id"`
	Name          string       `gorm:"type:varchar(100);not null" json:"cst_name"`
	DOB           time.Time    `gorm:"type:date;not null" json:"cst_dob" time_format:"2006-01-02"`
	PhoneNumber   string       `gorm:"type:varchar(20);not null" json:"cst_phone_number"`
	Email         string       `gorm:"type:varchar(100);not null;uniqueIndex" json:"cst_email"`
	NationalityID uint         `gorm:"not null;index" json:"nationality_id"`
	Nationality   Nationality  `gorm:"foreignKey:NationalityID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"nationality,omitempty"`
	FamilyLists   []FamilyList `gorm:"foreignKey:CustomerID" json:"family_list,omitempty"`
	CreatedAt     time.Time    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Customer) TableName() string { return "customer" }
