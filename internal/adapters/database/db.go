package database

import (
	"fmt"

	"github.com/billowdev/document-system-field-manager/internal/adapters/database/models"
	"github.com/billowdev/document-system-field-manager/pkg/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase() (*gorm.DB, error) {
	if configs.DB_SCHEMA == "" {
		configs.DB_SCHEMA = "public"
	}

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v search_path=%v",
		configs.DB_HOST,
		configs.DB_USERNAME,
		configs.DB_PASSWORD,
		configs.DB_NAME,
		configs.DB_PORT,
		configs.DB_SSL_MODE,
		configs.DB_SCHEMA,
	)
	loggerDBLevel := logger.Silent
	if configs.APP_DEBUG_MODE {
		loggerDBLevel = logger.Info
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		// PreferSimpleProtocol: DB_DRY_RUN,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(loggerDBLevel), // or logger.Silent if you don't want logs
		// Logger: logger.Default.LogMode(logger.Info), // or logger.Silent if you don't want logs
	})

	if err != nil {
		return nil, err // instead of panic, return the error
	}
	// !: ENABLE MIGRATIONS DB
	if configs.DB_RUN_MIGRATION {
		if err := RunMigrations(db); err != nil {
			return nil, err
		}
	}
	if configs.DB_RUN_SEEDER {
		RunSeeds(db)
	}
	return db, nil
}

func RunSeeds(db *gorm.DB) {

}

func RunMigrations(db *gorm.DB) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error
		if err != nil {
			return err
		}
		err = tx.AutoMigrate(
			// TODO START USER
			&models.Order{},
			&models.Document{},
			&models.SystemField{},
			&models.SystemGroupField{},
			&models.MasterFile{},
			&models.LogMasterFile{},
			&models.ConfigSystemMasterFileField{},
			&models.DocumentTemplate{},
			&models.DocumentVersion{},
			&models.DocumentTemplateField{},
			&models.DocumentVersionFieldValue{},
			&models.LogDocumentVersionFieldValue{},
		)
		return err
	})

	return err
}
