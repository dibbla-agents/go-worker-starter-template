// Package jobs provides job orchestration for multi-step workflows.
//
// This file contains a SIMPLE job example.
// For a more comprehensive example, see example_job.go
package jobs

import (
	"fmt"
	"log"
	"time"
)

// SimpleJob demonstrates a minimal job structure.
// Jobs orchestrate multi-step operations with error handling and metrics.
type SimpleJob struct {
	Name  string
	Count int
}

// SimpleJobResult contains the job execution results.
type SimpleJobResult struct {
	ItemsProcessed int
	Success        bool
	Error          error
	Duration       time.Duration
}

// NewSimpleJob creates a new simple job instance.
func NewSimpleJob(name string, count int) *SimpleJob {
	return &SimpleJob{
		Name:  name,
		Count: count,
	}
}

// Execute runs the job and returns the result.
func (j *SimpleJob) Execute() *SimpleJobResult {
	start := time.Now()
	result := &SimpleJobResult{}

	log.Printf("🚀 Starting job: %s", j.Name)

	// Step 1: Validate
	if j.Count <= 0 {
		result.Error = fmt.Errorf("count must be positive")
		result.Duration = time.Since(start)
		log.Printf("❌ Job failed: %v", result.Error)
		return result
	}

	// Step 2: Process
	log.Printf("   Processing %d items...", j.Count)
	for i := 0; i < j.Count; i++ {
		// Simulate work
		time.Sleep(10 * time.Millisecond)
		result.ItemsProcessed++
	}

	// Success
	result.Success = true
	result.Duration = time.Since(start)
	log.Printf("✅ Job completed: processed %d items in %s", result.ItemsProcessed, result.Duration)

	return result
}

