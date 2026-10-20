package bootstrap

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"clean-template/db/postgres"

	"github.com/golang-migrate/migrate/v4"
	pgmigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	pgdriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDatabase(env *Env) postgres.PostgresDatabase {
	dbHost := env.Postgres.DBHost
	dbPort := env.Postgres.DBPort
	dbUser := env.Postgres.DBUser
	dbPass := env.Postgres.DBPass
	dbName := env.Postgres.DBName
	dbAdmin := env.Postgres.DBAdmin
	sbAdminPass := env.Postgres.DBAdminPass

	// Connect to the PostgreSQL server first (without specifying a database).
	dsnWithoutDB := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable",
		dbHost, dbPort, dbAdmin, sbAdminPass)
	dbWithoutDB, err := gorm.Open(pgdriver.Open(dsnWithoutDB), &gorm.Config{})
	if err != nil {
		panic("failed to connect to PostgreSQL server: " + err.Error())
	}

	// Create the database if it does not exist.
	createDBSQL := fmt.Sprintf("CREATE DATABASE %s", dbName)
	dbWithoutDB.Exec(createDBSQL) // Ignore the error: the database may already exist.
	// Close the server-level connection.
	sqlDBWithoutDB, _ := dbWithoutDB.DB()
	sqlDBWithoutDB.Close()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s  sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)
	db, err := gorm.Open(pgdriver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to PostgreSQL: " + err.Error())
	}
	//// Run database migrations.
	//if err := runMigrations(dsn); err != nil {
	//	panic("failed to run migrations: " + err.Error())
	//}

	// Create the database schema.
	if err := CreateDatabase(db); err != nil {
		//panic("failed to create database: " + err.Error())
		slog.Warn("failed to create database", "err", err)
	}

	return postgres.PostgresDatabase{DB: db}
}

// runMigrations applies the .sql files under db/migrations.
func runMigrations(dsn string) error {
	driver, err := pgmigrate.WithInstance(nil, &pgmigrate.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %w", err)
	}

	// Make sure the migration source path is correct.
	migrationPath := "file://db/migrations"
	absPath, err := filepath.Abs("db/migrations")
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return fmt.Errorf("migrations directory does not exist: %s", absPath)
	}

	m, err := migrate.NewWithDatabaseInstance(migrationPath, "postgres", driver)
	if err != nil {
		return fmt.Errorf("failed to initialize migration: %w", err)
	}

	// Apply all pending migrations.
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	return nil
}

// create database using gorm with .sql file
func CreateDatabase(db *gorm.DB) error {
	// Read the bootstrap SQL file.
	sqlBytes, err := os.ReadFile("db/postgres/migrations/file_management.sql")
	if err != nil {
		Fatal("failed to read SQL file", "err", err)
	}
	sql := string(sqlBytes)

	err = db.Exec(sql).Error
	if err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}
	return nil
}

// ClosePostgresConnection closes the PostgreSQL connection.
func ClosePostgresConnection(db postgres.PostgresDatabase) {
	sqlDB, err := db.DB.DB()
	if err != nil {
		slog.Error("failed to get sql.DB", "err", err)
		return
	}
	if err := sqlDB.Close(); err != nil {
		slog.Error("failed to close PostgreSQL connection", "err", err)
	}
	slog.Info("Connection to Postgres closed.")
}
