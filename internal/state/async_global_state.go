// Package state provides centralized state management for worker functions.
//
// STARTER TEMPLATE INSTRUCTIONS:
// This file demonstrates the AsyncGlobalState pattern for sharing resources across
// your worker functions (database connections, API clients, caches, etc.)
//
// To customize for your application:
// 1. Uncomment the fields in AsyncGlobalState that you need
// 2. Uncomment the initialization code in NewAsyncGlobalState()
// 3. Update the Close() method if you have resources to clean up
// 4. Call NewAsyncGlobalState() once in your main.go when the application starts
// 5. Pass the AsyncGlobalState instance to your worker functions
//
// If you don't need shared state, you can delete this entire package.
package state

import (
	// Uncomment if using GORM with PostgreSQL
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"

	// Uncomment if you have models to migrate
	// "github.com/dibbla-agents/go-worker-starter-template/internal/models"
)

// AsyncGlobalState provides centralized state management for worker functions.
// Use this to share resources across your application and prevent duplicate connections.
//
// Benefits:
// - Single database connection pool shared by all workers
// - Centralized API client instances
// - Shared cache connections
// - Configuration access
type AsyncGlobalState struct {
	// TODO: Uncomment and add the shared resources you need
	
	// Database connection (shared singleton)
	// Prevents creating multiple connections for each worker function
	// DB *gorm.DB
	
	// API clients
	// ExternalAPIClient *api.Client
	
	// Cache connections
	// RedisClient *redis.Client
	
	// Configuration
	// Config *config.Config
	
	// Metrics and monitoring
	// MetricsCollector *metrics.Collector
}

// NewAsyncGlobalState initializes the global state with all shared resources.
// Call this once when your application starts (typically in cmd/worker/main.go).
//
// Example usage in main.go:
//
//   ags, err := state.NewAsyncGlobalState()
//   if err != nil {
//       log.Fatalf("Failed to initialize global state: %v", err)
//   }
//   defer ags.Close()
//
func NewAsyncGlobalState() (*AsyncGlobalState, error) {
	// TODO: Uncomment and initialize the resources you need
	
	// Example: Database initialization
	// Read database URL from environment
	// dsn := os.Getenv("DATABASE_URL")
	// if dsn == "" {
	// 	return nil, fmt.Errorf("DATABASE_URL environment variable not set")
	// }
	
	// Connect to PostgreSQL using GORM
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
	// 	// Optional: Configure connection pool settings
	// 	// PrepareStmt: true,
	// 	// DisableForeignKeyConstraintWhenMigrating: true,
	// })
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to connect to database: %w", err)
	// }
	
	// Auto-migrate database tables (creates/updates schema)
	// Note: In production, consider using proper migration tools instead of AutoMigrate
	// if err := db.AutoMigrate(models.AllModels()...); err != nil {
	// 	return nil, fmt.Errorf("failed to migrate database tables: %w", err)
	// }
	
	// Example: Initialize API client
	// cfg, err := config.Load()
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to load config: %w", err)
	// }
	// apiClient := api.NewClient(cfg.APIKey, cfg.APIBaseURL)
	
	// Example: Initialize Redis cache
	// redisClient := redis.NewClient(&redis.Options{
	// 	Addr: os.Getenv("REDIS_URL"),
	// })
	
	return &AsyncGlobalState{
		// TODO: Uncomment and assign your initialized resources
		// DB: db,
		// ExternalAPIClient: apiClient,
		// RedisClient: redisClient,
		// Config: cfg,
	}, nil
}

// Close gracefully closes all connections and releases resources.
// Call this during application shutdown (typically with defer in main.go).
func (ags *AsyncGlobalState) Close() error {
	// TODO: Uncomment and add cleanup for your resources
	
	// Example: Close database connection
	// if ags.DB != nil {
	// 	sqlDB, err := ags.DB.DB()
	// 	if err != nil {
	// 		return fmt.Errorf("failed to get underlying SQL DB: %w", err)
	// 	}
	// 	if err := sqlDB.Close(); err != nil {
	// 		return fmt.Errorf("failed to close database: %w", err)
	// 	}
	// }
	
	// Example: Close Redis connection
	// if ags.RedisClient != nil {
	// 	if err := ags.RedisClient.Close(); err != nil {
	// 		return fmt.Errorf("failed to close Redis: %w", err)
	// 	}
	// }
	
	return nil
}

// Example worker function signature using AsyncGlobalState:
//
// func MyWorkerFunction(ctx context.Context, args MyArgs, ags *state.AsyncGlobalState) error {
//     // Access shared resources from ags
//     var result MyModel
//     err := ags.DB.Where("id = ?", args.ID).First(&result).Error
//     if err != nil {
//         return fmt.Errorf("failed to query database: %w", err)
//     }
//     return nil
// }

