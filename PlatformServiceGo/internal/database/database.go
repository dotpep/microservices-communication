package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/dotpep/microservices-communication/PlatformServiceGo/internal/models"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error

	// GORM Related
	GetDB() *gorm.DB
	RunMigrations()
}

type service struct {
	db    *gorm.DB
	sqlDB *sql.DB
}

var (
	database   = os.Getenv("POSTGRES_DB_NAME")
	password   = os.Getenv("POSTGRES_DB_PASSWORD")
	username   = os.Getenv("POSTGRES_DB_USERNAME")
	port       = os.Getenv("POSTGRES_DB_PORT_API")
	host       = os.Getenv("POSTGRES_DB_HOST")
	schema     = os.Getenv("POSTGRES_DB_SCHEMA")
	dbInstance *service
)

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

	// Build connection string
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		username, password, host, port, database, schema,
	)

	gormDB, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// TODO: Fix crutches in this code, cut sqlDb and also Health() function implementation with GORM, to check database connection health
	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB: %v", err)
	}

	// Setup database instance settings
	dbInstance = &service{
		db:    gormDB,
		sqlDB: sqlDB,
	}

	return dbInstance
}

// RunMigrations run auto migration for given models of GORM
func (s *service) RunMigrations() {
	log.Println("---> Running database migrations (GORM Auto)...")
	if err := s.db.AutoMigrate(
		// Add/List models here:
		&models.Platform{},
	); err != nil {
		log.Fatalf("Failed to run migrations: %v\n", err)
	}
	log.Println("---> Database migrations completed.")
}

// GetDB returns the underlying GORM DB instance
func (s *service) GetDB() *gorm.DB {
	return s.db
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the database
	err := s.sqlDB.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf("db down: %v", err) // Log the error and terminate the program
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := s.sqlDB.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}

	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", database)
	return s.sqlDB.Close()
}
