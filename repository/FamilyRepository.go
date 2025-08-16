package repository

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"test_booking/entity"
)

type FamilyRepository interface {
	Save(ctx context.Context, family entity.FamilyList) (entity.FamilyList, error)
	Update(ctx context.Context, id uint, updateFamily entity.FamilyList) (entity.FamilyList, error)
	Delete(ctx context.Context, id uint) (string, error)
	GetFamilyByCustID(ctx context.Context, castId uint) ([]entity.FamilyList, error)
}
type familyConnection struct {
	connection *gorm.DB
}

func (c familyConnection) GetFamilyByCustID(ctx context.Context, castId uint) ([]entity.FamilyList, error) {
	//TODO implement me
	var family []entity.FamilyList

	err := c.connection.WithContext(ctx).
		Find(&family, "customer_id = ?", castId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return family, fmt.Errorf("data family not found")
		}
		return family, fmt.Errorf("failed to get family: %w", err)
	}

	return family, nil
}

func (c familyConnection) Delete(ctx context.Context, id uint) (string, error) {
	tx := c.connection.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var family entity.FamilyList
	if err := tx.First(&family, id).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", fmt.Errorf("family not found")
		}
		return "", fmt.Errorf("failed to find family: %w", err)
	}

	if err := tx.Delete(&family).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to delete family: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	return "family deleted successfully", nil
}

func (c familyConnection) Save(ctx context.Context, family entity.FamilyList) (entity.FamilyList, error) {
	tx := c.connection.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&family).Error; err != nil {
		tx.Rollback()
		return family, fmt.Errorf("failed to create family: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return family, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return family, nil
}

func (c familyConnection) Update(ctx context.Context, id uint, updateFamily entity.FamilyList) (entity.FamilyList, error) {
	tx := c.connection.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 1. Ambil data existing
	var exisitingFamily entity.FamilyList
	if err := tx.First(&exisitingFamily, id).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.FamilyList{}, fmt.Errorf("family not found")
		}
		return entity.FamilyList{}, fmt.Errorf("failed to get family: %w", err)
	}

	// 2. Update data
	if err := tx.Model(&exisitingFamily).Updates(updateFamily).Error; err != nil {
		tx.Rollback()
		if strings.Contains(err.Error(), "duplicate key") {
			return entity.FamilyList{}, fmt.Errorf("email already exists")
		}

		return entity.FamilyList{}, fmt.Errorf("failed to update family: %w", err)
	}

	var result entity.FamilyList

	if err := tx.Commit().Error; err != nil {
		return entity.FamilyList{}, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return result, nil
}

func NewFamilyRepository(db *gorm.DB) FamilyRepository {
	return &familyConnection{
		connection: db,
	}
}
