package tasks

import (
	"fmt"
	"time"

	"github.com/dibbla-agents/go-worker-starter-template/internal/state"
)

// ExampleTask demonstrates the task pattern.
// Tasks are reusable, single-purpose operations that jobs orchestrate.
type ExampleTask struct {
	AGS    *state.AsyncGlobalState
	Limit  int    // Task configuration
	Filter string // Additional parameters
}

// ExampleTaskResult contains the task execution results
type ExampleTaskResult struct {
	ItemsProcessed int
	ProcessedData  []string
	ExecutionTime  time.Duration
}

// Execute runs the task and returns results
func (t *ExampleTask) Execute() (*ExampleTaskResult, error) {
	start := time.Now()

	// 1. Validate inputs
	if err := t.validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// 2. Perform the task's core operation
	data, err := t.process()
	if err != nil {
		return nil, fmt.Errorf("processing failed: %w", err)
	}

	// 3. Return results
	return &ExampleTaskResult{
		ItemsProcessed: len(data),
		ProcessedData:  data,
		ExecutionTime:  time.Since(start),
	}, nil
}

// validate checks task preconditions
func (t *ExampleTask) validate() error {
	if t.Limit <= 0 {
		return fmt.Errorf("limit must be positive")
	}
	return nil
}

// process implements the task's main logic
func (t *ExampleTask) process() ([]string, error) {
	results := make([]string, 0)

	// Example: Query database
	// if t.AGS != nil && t.AGS.DB != nil {
	//     var records []models.YourModel
	//     query := t.AGS.DB.Limit(t.Limit)
	//     if t.Filter != "" {
	//         query = query.Where("field = ?", t.Filter)
	//     }
	//     if err := query.Find(&records).Error; err != nil {
	//         return nil, fmt.Errorf("database query failed: %w", err)
	//     }
	//     
	//     for _, record := range records {
	//         results = append(results, record.Name)
	//     }
	// }

	// Simulate processing
	for i := 0; i < t.Limit; i++ {
		results = append(results, fmt.Sprintf("item_%d", i))
	}

	return results, nil
}

// NewExampleTask creates a new task with default configuration
func NewExampleTask(ags *state.AsyncGlobalState, limit int) *ExampleTask {
	return &ExampleTask{
		AGS:    ags,
		Limit:  limit,
		Filter: "",
	}
}

// NewExampleTaskWithFilter creates a new task with filtering
func NewExampleTaskWithFilter(ags *state.AsyncGlobalState, limit int, filter string) *ExampleTask {
	return &ExampleTask{
		AGS:    ags,
		Limit:  limit,
		Filter: filter,
	}
}

