package database

import (
	"jwt-gin-gorm/internals/user/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	DB *gorm.DB
}

func NewDatabase(connectionString string) (*PostgresDB, error) {

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	// configure connection pooling
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	// DB migration
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}

	return &PostgresDB{DB: db}, nil

}

func (p *PostgresDB) Close() error {
	sqlDB, err := p.DB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
