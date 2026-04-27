package db

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Query is the global GORM instance.
var Query *gorm.DB

func setCreatedAt(db *gorm.DB) {
	if db.Statement.Schema != nil {
		createdAtField := db.Statement.Schema.LookUpField("CreatedAt")
		if createdAtField != nil {
			reflectValue := db.Statement.ReflectValue
			if reflectValue.Kind() == reflect.Slice {
				for i := 0; i < reflectValue.Len(); i++ {
					elem := reflectValue.Index(i)
					if elem.Kind() == reflect.Ptr {
						elem = elem.Elem()
					}
					_, isZero := createdAtField.ValueOf(db.Statement.Context, elem)
					if isZero {
						_ = createdAtField.Set(db.Statement.Context, elem, time.Now())
					}
				}
			} else {
				_, isZero := createdAtField.ValueOf(db.Statement.Context, reflectValue)
				if isZero {
					_ = createdAtField.Set(db.Statement.Context, reflectValue, time.Now())
				}
			}
		}
	}
}

func setUpdatedAt(db *gorm.DB) {
	if db.Statement.Schema != nil {
		field := db.Statement.Schema.LookUpField("UpdatedAt")
		if field != nil {
			_ = field.Set(db.Statement.Context, db.Statement.ReflectValue, time.Now())
		}
	}
}

// EnableGlobalTimestamps enables auto timestamps.
func EnableGlobalTimestamps(db *gorm.DB) {
	db.Callback().Create().Before("gorm:create").Register("global:set_created_at", setCreatedAt)
	db.Callback().Update().Before("gorm:update").Register("global:set_updated_at", setUpdatedAt)
}

// OpenDBConnection returns a DB connection.
func OpenDBConnection() (*gorm.DB, error) {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using system env")
	}

	maxConnStr := os.Getenv("DB_MAX_CONNECTIONS")
	maxIdleConnStr := os.Getenv("DB_MAX_IDLE_CONNECTIONS")
	maxLifetimeConnStr := os.Getenv("DB_MAX_LIFETIME_CONNECTIONS")

	maxConn, err := strconv.Atoi(maxConnStr)
	if err != nil || maxConn == 0 {
		maxConn = 10
	}

	maxIdleConn, err := strconv.Atoi(maxIdleConnStr)
	if err != nil || maxIdleConn == 0 {
		maxIdleConn = 5
	}

	maxLifetimeConn, err := strconv.Atoi(maxLifetimeConnStr)
	if err != nil || maxLifetimeConn == 0 {
		maxLifetimeConn = 300
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DBNAME"),
		os.Getenv("DB_SSLMODE"),
	)

	Query, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connected:", Query.Name())

	sqlDB, sqlDBErr := Query.DB()
	if sqlDBErr != nil {
		log.Println("Failed to get sql.DB:", sqlDBErr)
	}

	sqlDB.SetMaxOpenConns(maxConn)
	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetConnMaxLifetime(time.Duration(maxLifetimeConn) * time.Second)

	return Query, nil
}
