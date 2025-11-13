// Package models defines database models for GORM.
//
// STARTER TEMPLATE INSTRUCTIONS:
// This file demonstrates common GORM model patterns and table creation.
//
// To customize for your application:
// 1. Uncomment and modify the example models below
// 2. Add your models to the AllModels() function
// 3. GORM will automatically create/update tables when AutoMigrate is called
// 4. Review GORM tags for field constraints and relationships
//
// If you're not using a database, you can delete this entire package.
package models

import (
	"time"
	// Uncomment if using GORM
	// "gorm.io/gorm"
)

// AllModels returns all models to be registered with GORM AutoMigrate.
// Add any new models to this list to automatically create/update their tables.
//
// Usage in state/async_global_state.go:
//   db.AutoMigrate(models.AllModels()...)
func AllModels() []interface{} {
	return []interface{}{
		// TODO: Uncomment and add your models here
		// &User{},
		// &Task{},
		// &TaskResult{},
	}
}

// Example Model 1: User
// Demonstrates basic model with common fields and GORM tags
//
// type User struct {
// 	// Primary key (auto-increment)
// 	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
// 	
// 	// Unique fields
// 	Email    string `gorm:"uniqueIndex;not null" json:"email"`
// 	Username string `gorm:"uniqueIndex;not null" json:"username"`
// 	
// 	// Regular fields
// 	Name     string `gorm:"size:255" json:"name"`
// 	Age      int    `gorm:"default:0" json:"age"`
// 	IsActive bool   `gorm:"default:true" json:"is_active"`
// 	
// 	// Relationships
// 	Tasks []Task `gorm:"foreignKey:UserID" json:"tasks,omitempty"`
// 	
// 	// Timestamps (automatically managed by GORM)
// 	CreatedAt time.Time `json:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at"`
// 	
// 	// Soft delete (if deleted, DeletedAt is set instead of actual deletion)
// 	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
// }

// Example Model 2: Task
// Demonstrates relationships and foreign keys
//
// type Task struct {
// 	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
// 	Title       string `gorm:"not null;size:255" json:"title"`
// 	Description string `gorm:"type:text" json:"description"`
// 	Status      string `gorm:"type:varchar(50);default:'pending';index" json:"status"` // pending, processing, completed, failed
// 	
// 	// Foreign key relationship
// 	UserID uint  `gorm:"not null;index" json:"user_id"`
// 	User   *User `gorm:"constraint:OnDelete:CASCADE" json:"user,omitempty"`
// 	
// 	// Has many relationship
// 	Results []TaskResult `gorm:"foreignKey:TaskID" json:"results,omitempty"`
// 	
// 	// Timestamps
// 	CreatedAt   time.Time  `json:"created_at"`
// 	UpdatedAt   time.Time  `json:"updated_at"`
// 	CompletedAt *time.Time `json:"completed_at,omitempty"` // Nullable timestamp
// }

// Example Model 3: TaskResult
// Demonstrates storing worker function results
//
// type TaskResult struct {
// 	ID     uint   `gorm:"primaryKey;autoIncrement" json:"id"`
// 	TaskID uint   `gorm:"not null;index" json:"task_id"`
// 	Task   *Task  `gorm:"constraint:OnDelete:CASCADE" json:"task,omitempty"`
// 	
// 	// Result data
// 	Output    string `gorm:"type:text" json:"output"`      // Success output
// 	Error     string `gorm:"type:text" json:"error"`       // Error message if failed
// 	Success   bool   `gorm:"default:false" json:"success"` // Whether task succeeded
// 	Duration  int64  `gorm:"default:0" json:"duration"`    // Execution time in milliseconds
// 	
// 	// Metadata
// 	WorkerID string `gorm:"size:255;index" json:"worker_id"` // Which worker processed it
// 	
// 	CreatedAt time.Time `json:"created_at"`
// }

// Common GORM Tags Reference:
//
// Field Constraints:
//   - primaryKey         : Mark as primary key
//   - autoIncrement      : Auto-incrementing field
//   - unique             : Unique constraint
//   - uniqueIndex        : Create unique index
//   - not null           : NOT NULL constraint
//   - default:value      : Default value
//   - size:255           : VARCHAR size
//   - type:text          : Column type override
//
// Indexes:
//   - index              : Create index
//   - index:idx_name     : Named index
//   - uniqueIndex        : Unique index
//
// Relationships:
//   - foreignKey:Field   : Specify foreign key
//   - references:ID      : Reference field
//   - constraint:OnDelete:CASCADE : Delete cascade
//
// Other:
//   - ->                 : Read only
//   - -                  : Ignore field
//   - embedded           : Embed struct
//
// JSON Tags:
//   - json:"field_name"  : JSON field name
//   - json:",omitempty"  : Omit if empty
//   - json:"-"           : Don't serialize to JSON
//
// Example Custom Table Name:
//
// func (User) TableName() string {
//     return "custom_users_table"
// }

