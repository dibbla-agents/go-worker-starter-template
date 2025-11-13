// Package config provides application configuration management.
//
// STARTER TEMPLATE INSTRUCTIONS:
// This file demonstrates a common pattern for managing application configuration
// using environment variables with sensible defaults.
//
// To customize for your application:
// 1. Uncomment the fields in the Config struct that you need
// 2. Uncomment the corresponding lines in Load() function
// 3. Update Validate() to check your required fields
// 4. Create a .env file in your project root with the required values
// 5. Update environment variable names and defaults as needed
//
// If you don't need configuration management, you can delete this entire package.
package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application.
// TODO: Uncomment and customize fields for your specific application needs.
type Config struct {
	// External API configuration
	// Uncomment if you're integrating with external APIs
	// APIKey     string
	// APIBaseURL string

	// Database configuration
	// Uncomment if you're using a database
	// DatabaseURL string

	// Worker configuration
	// Uncomment if you need worker service settings
	// WorkerHost string
	// WorkerPort string

	// Application settings
	// Uncomment the settings you need
	// AppEnv   string // Environment: development, staging, production
	// LogLevel string // Logging level: debug, info, warn, error
	// DataDir  string // Directory for application data files
}

// Load reads configuration from environment variables and .env file.
// It returns a validated Config struct or an error if required values are missing.
func Load() (*Config, error) {
	// Load .env file if it exists (optional - fails silently)
	// In production, you typically use actual environment variables instead
	_ = godotenv.Load()

	cfg := &Config{
		// TODO: Uncomment and configure the fields you need
		
		// External API configuration
		// APIKey:      os.Getenv("API_KEY"),
		// APIBaseURL:  getEnvOrDefault("API_BASE_URL", "https://api.example.com/v1"),
		
		// Database configuration
		// DatabaseURL: os.Getenv("DATABASE_URL"),
		
		// Worker configuration
		// WorkerHost:  getEnvOrDefault("WORKER_HOST", "localhost"),
		// WorkerPort:  getEnvOrDefault("WORKER_PORT", "50051"),
		
		// Application settings
		// AppEnv:      getEnvOrDefault("APP_ENV", "development"),
		// LogLevel:    getEnvOrDefault("LOG_LEVEL", "info"),
		// DataDir:     getEnvOrDefault("DATA_DIR", "./data"),
	}

	// Validate the configuration before returning
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

// Validate checks that all required configuration values are present.
// TODO: Uncomment and add validation rules based on your required fields.
func (c *Config) Validate() error {
	// Example validations - uncomment and modify based on your requirements
	
	// if c.APIKey == "" {
	// 	return fmt.Errorf("API_KEY is required")
	// }
	
	// if c.DatabaseURL == "" {
	// 	return fmt.Errorf("DATABASE_URL is required")
	// }
	
	// if c.WorkerHost == "" {
	// 	return fmt.Errorf("WORKER_HOST is required")
	// }
	
	return nil
}

// getEnvOrDefault retrieves an environment variable or returns a default value.
// This is a helper function to simplify config initialization.
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// Example usage:
//
// In your main.go or other initialization code:
//
//   cfg, err := config.Load()
//   if err != nil {
//       log.Fatalf("Failed to load config: %v", err)
//   }
//
//   // Use config values throughout your application
//   client := api.NewClient(cfg.APIKey, cfg.APIBaseURL)
