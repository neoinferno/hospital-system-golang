package database

import (
	"fmt"
	"time"

	"hospital-system/configs"
	"hospital-system/internal/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Service interface {
	DB() *gorm.DB
	Close() error
}

type service struct {
	db *gorm.DB
}

var (
	database   string
	password   string
	username   string
	port       string
	host       string
	dbInstance *service
)

func init() {
	secretCfg := configs.GetSecret()
	database = secretCfg.Postgres.Db
	password = secretCfg.Postgres.Password
	username = secretCfg.Postgres.User
	port = secretCfg.Postgres.Port
	host = secretCfg.Postgres.Host
}

func NewService() Service {
	if dbInstance != nil {
		return dbInstance
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, username, password, database, "disable")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		QueryFields: true,
		Logger:      GormLogger{logger.Default.LogMode(logger.Info)},
	})
	if err != nil {
		fmt.Printf("failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(
		entities.Patient{},
		entities.Staff{},
		entities.Hospital{},
	)
	if err != nil {
		fmt.Printf("Failed to migrate database: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("Failed to get database connection: %v", err)
	}
	if err == nil {
		sqlDB, _ := db.DB()
		sqlDB.Exec("SET TIME ZONE 'UTC';")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("Postgres initialized")
	dbInstance = &service{db: db}
	return dbInstance
}

func (s *service) DB() *gorm.DB {
	return s.db
}

func (s *service) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB instance: %w", err)
	}
	return sqlDB.Close()
}

type GormLogger struct {
	logger.Interface
}
