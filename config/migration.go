package config

import (
	"fmt"

	"gorm.io/gorm"
	"test_booking/entity"
)

func Migration(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.AutoMigrate(
			&entity.Nationality{},
			&entity.Customer{},
			&entity.FamilyList{},
		); err != nil {
			return fmt.Errorf("automigrate: %w", err)
		}
		return nil
	})
}
