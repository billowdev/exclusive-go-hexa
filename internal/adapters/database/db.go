package database

import (
	"fmt"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/models"
	"github.com/billowdev/exclusive-go-hexa/internal/adapters/database/seeders"
	"github.com/billowdev/exclusive-go-hexa/pkg/configs"
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
	_ = seeders.SEED_ORDER
	seeders.SeedOrder(db)
	seeders.SeedSystemField(db)
	seeders.SeedGroupField(db)
	seeders.SeedConfigSystemMasterFileField(db)
	seeders.SeedMasterFile(db)
	seeders.SeedLogMasterFile(db)
	seeders.SeedDocument(db)
	seeders.SeedDocumentTemplate(db)
	seeders.SeedDocumentTemplateField(db)
	seeders.SeedDocumentVersion(db)
	seeders.SeedDocumentVersionFieldValue(db)
	seeders.SeedLogDocumentVersionFieldValue(db)
}

func resetSeeder(db *gorm.DB) error {
	if err := helperDeleteInfo(db, models.TNConfigSystemMasterFileField); err != nil {
		return err
	}

	if err := helperDeleteInfo(db, models.TNLogMasterFile); err != nil {
		return err
	}

	if err := helperDeleteInfo(db, models.TNLogDocumentVersionFieldValue); err != nil {
		return err
	}

	if err := helperDeleteInfo(db, models.TNMasterFile); err != nil {
		return err
	}
	if err := helperDeleteInfo(db, models.TNDocumentVersionFieldValue); err != nil {
		return err
	}

	if err := helperDeleteInfo(db, models.TNDocumentVersion); err != nil {
		return err
	}

	if err := helperDeleteInfo(db, models.TNDocumentTemplateField); err != nil {
		return err
	}

	if err := helperDeleteInfo(db, models.TNDocumentTemplate); err != nil {
		return err
	}

	if err := helperDeleteInfo(db, models.TNDocument); err != nil {
		return err
	}
	if err := helperDeleteInfo(db, models.TNOrder); err != nil {
		return err
	}

	if err := helperDeleteInfo(db, models.TNSystemField); err != nil {
		return err
	}

	if err := helperDeleteInfo(db, models.TNSystemGroupField); err != nil {
		return err
	}
	return nil
}
func helperDeleteInfo(db *gorm.DB, table string) error {
	err := db.Exec(fmt.Sprintf("DELETE FROM %s", table)).Error
	if err != nil {
		return err
	}
	err = db.Exec(fmt.Sprintf("SELECT setval('%s_id_seq', 1, false)", table)).Error
	if err != nil {
		return err
	}
	return nil
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
		if err != nil {
			return err
		}
		err = resetSeeder(db)
		if err != nil {
			return err
		}
		return err
	})

	return err
}
