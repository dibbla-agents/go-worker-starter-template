package jobs

import (
	"fmt"
	"log"
	"time"

	"worker_starter_template/internal/jobs/tasks"
	"worker_starter_template/internal/state"
)

// ExampleJob demonstrates the standard job pattern.
// Jobs orchestrate multiple tasks and handle error recovery, logging, and metrics.
type ExampleJob struct {
	AGS       *state.AsyncGlobalState
	ItemCount int    // Job configuration parameter
	JobType   string // Identifier for this job
}

// ExampleJobResult contains execution results and statistics
type ExampleJobResult struct {
	ItemsProcessed int
	Success        bool
	Error          error
	ExecutionTime  time.Duration
}

// Execute runs the job
func (j *ExampleJob) Execute() *ExampleJobResult {
	result := &ExampleJobResult{
		Success: false,
	}

	start := time.Now()

	log.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	log.Printf("🚀 Starting %s Job", j.JobType)
	log.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	// Phase 1: Validate inputs
	if err := j.validate(); err != nil {
		result.Error = fmt.Errorf("validation failed: %w", err)
		result.ExecutionTime = time.Since(start)
		j.logError(result)
		return result
	}

	// Phase 2: Execute main logic (using tasks)
	log.Println("📋 Processing items...")
	processed, err := j.process()
	if err != nil {
		result.Error = fmt.Errorf("processing failed: %w", err)
		result.ExecutionTime = time.Since(start)
		j.logError(result)
		return result
	}

	result.ItemsProcessed = processed
	log.Printf("   ✓ Processed %d items", processed)

	// Job completed successfully
	result.Success = true
	result.ExecutionTime = time.Since(start)
	j.logSuccess(result)

	return result
}

// validate checks job preconditions
func (j *ExampleJob) validate() error {
	if j.ItemCount <= 0 {
		return fmt.Errorf("item count must be positive")
	}
	return nil
}

// process implements the main job logic
func (j *ExampleJob) process() (int, error) {
	// Example 1: Simple inline logic
	// time.Sleep(100 * time.Millisecond)
	// return j.ItemCount, nil

	// Example 2: Using a task (recommended for complex operations)
	task := tasks.NewExampleTask(j.AGS, j.ItemCount)
	result, err := task.Execute()
	if err != nil {
		return 0, fmt.Errorf("task execution failed: %w", err)
	}

	log.Printf("   ✓ Task completed in %s", result.ExecutionTime)
	return result.ItemsProcessed, nil
}

// logSuccess prints success summary
func (j *ExampleJob) logSuccess(result *ExampleJobResult) {
	log.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	log.Println("✅ Job Completed Successfully")
	log.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	log.Printf("📊 Items processed: %d", result.ItemsProcessed)
	log.Printf("⏱️  Execution time: %s", result.ExecutionTime.Round(time.Millisecond))
	log.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
}

// logError prints error summary
func (j *ExampleJob) logError(result *ExampleJobResult) {
	log.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	log.Println("❌ Job Failed")
	log.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	log.Printf("💥 Error: %v", result.Error)
	log.Printf("📊 Items processed: %d", result.ItemsProcessed)
	log.Printf("⏱️  Time elapsed: %s", result.ExecutionTime.Round(time.Millisecond))
	log.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
}

// NewExampleJob creates a new job instance
func NewExampleJob(ags *state.AsyncGlobalState, itemCount int) *ExampleJob {
	return &ExampleJob{
		AGS:       ags,
		ItemCount: itemCount,
		JobType:   "example",
	}
}

