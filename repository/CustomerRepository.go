package repository

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"test_booking/entity"
)

type CustomerRepository interface {
	Save(ctx context.Context, customer entity.Customer) (entity.Customer, error)
	GetCustomerByEmail(ctx context.Context, email string) (entity.Customer, error)
	Update(ctx context.Context, id uint, updatedCustomer entity.Customer) (entity.Customer, error)
	Detail(ctx context.Context, id uint) (entity.Customer, error)
	List(ctx context.Context) ([]entity.Customer, error)
	Delete(ctx context.Context, id uint) (string, error)
}
type customerConnection struct {
	connection *gorm.DB
}

func (c customerConnection) Delete(ctx context.Context, id uint) (string, error) {
	tx := c.connection.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var customer entity.Customer
	if err := tx.First(&customer, id).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", fmt.Errorf("customer not found")
		}
		return "", fmt.Errorf("failed to find customer: %w", err)
	}

	if err := tx.Delete(&customer).Error; err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to delete customer: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	return "customer deleted successfully", nil
}

func (c customerConnection) List(ctx context.Context) ([]entity.Customer, error) {
	var customers []entity.Customer

	if err := c.connection.WithContext(ctx).
		Preload("Nationality").
		Preload("FamilyLists").
		Find(&customers).Error; err != nil {
		return nil, fmt.Errorf("failed to list customers: %w", err)
	}

	return customers, nil
}

func (c customerConnection) Detail(ctx context.Context, id uint) (entity.Customer, error) {
	//TODO implement me
	var cust entity.Customer

	err := c.connection.WithContext(ctx).
		Preload("Nationality").
		Preload("FamilyLists").
		First(&cust, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Customer{}, fmt.Errorf("customer not found")
		}
		return entity.Customer{}, fmt.Errorf("failed to get customer: %w", err)
	}

	return cust, nil
}

func (c customerConnection) GetCustomerByEmail(ctx context.Context, email string) (entity.Customer, error) {
	var customer entity.Customer

	err := c.connection.WithContext(ctx).
		Where("email = ?", email).
		First(&customer).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		}
		return customer, err
	}

	return customer, nil
}

func (c customerConnection) Save(ctx context.Context, customer entity.Customer) (entity.Customer, error) {
	tx := c.connection.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&customer).Error; err != nil {
		tx.Rollback()
		return customer, fmt.Errorf("failed to create customer: %v", err)
	}

	var savedCustomer entity.Customer
	if err := tx.
		Preload("Nationality").
		Preload("FamilyLists").
		First(&savedCustomer, customer.ID).Error; err != nil {
		tx.Rollback()
		return customer, fmt.Errorf("failed to load saved customer: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return customer, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return savedCustomer, nil
}

func (c customerConnection) Update(ctx context.Context, id uint, updatedCustomer entity.Customer) (entity.Customer, error) {
	tx := c.connection.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 1. Ambil data existing
	var existingCustomer entity.Customer
	if err := tx.First(&existingCustomer, id).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Customer{}, fmt.Errorf("customer not found")
		}
		return entity.Customer{}, fmt.Errorf("failed to get customer: %w", err)
	}

	// 2. Update data
	if err := tx.Model(&existingCustomer).Updates(updatedCustomer).Error; err != nil {
		tx.Rollback()
		if strings.Contains(err.Error(), "duplicate key") {
			return entity.Customer{}, fmt.Errorf("email already exists")
		}

		return entity.Customer{}, fmt.Errorf("failed to update customer: %w", err)
	}

	var result entity.Customer
	if err := tx.
		Preload("Nationality").
		Preload("FamilyLists").
		First(&result, id).Error; err != nil {
		tx.Rollback()
		return entity.Customer{}, fmt.Errorf("failed to load updated customer: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return entity.Customer{}, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return result, nil
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerConnection{
		connection: db,
	}
}
