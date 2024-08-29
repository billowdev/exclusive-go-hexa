package database

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type TransactorImpl struct {
	db *gorm.DB
}

// BeginTransaction implements IDatabasePorts.
func (d *TransactorImpl) BeginTransaction() (*gorm.DB, error) {
	tx := d.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}

// RollbackTransaction rolls back the transaction if it was started and returns any error encountered.
func (d *TransactorImpl) RollbackTransaction(tx *gorm.DB) error {
	if tx == nil {
		return nil // No transaction to rollback
	}
	if tx.Error != nil {
		return tx.Error // If there was an error, return it
	}

	// Rollback the transaction
	if err := tx.Rollback().Error; err != nil {
		return err
	}
	return nil
}

// WithinTransaction implements IDatabaseTransactor.
func (d *TransactorImpl) WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error {
	// begin transaction
	tx, err := d.BeginTransaction()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", tx.Error)
	}

	// Ensure that the transaction is finalized properly
	defer func() {
		if r := recover(); r != nil {
			_ = d.RollbackTransaction(tx)
			panic(r) // Re-panic after rollback
		} else if tx.Error != nil {
			_ = d.RollbackTransaction(tx)
		} else {
			tx.Commit()
		}
	}()

	// Run the callback function with the transaction context
	err = tFunc(InjectTx(ctx, tx))
	if err != nil {
		tx.Error = err // Set the error to indicate a rollback is needed
		return err
	}

	return nil
}

type IDatabaseTransactor interface {
	// InjectTx(ctx context.Context, tx *gorm.DB) context.Context
	// ExtractTx(ctx context.Context) *gorm.DB
	WithinTransaction(context.Context, func(ctx context.Context) error) error
	BeginTransaction() (*gorm.DB, error)
	RollbackTransaction(tx *gorm.DB) error
}

func NewTransactorRepo(db *gorm.DB) IDatabaseTransactor {
	return &TransactorImpl{db: db}
}
