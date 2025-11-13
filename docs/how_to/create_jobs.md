# How to Create Jobs

Jobs orchestrate multi-step workflows with error handling, logging, and metrics. Unlike worker functions (which respond to external calls), jobs run internally on demand or on schedule.

## Job vs Worker Function

| Feature | Worker Function | Job |
|---------|----------------|-----|
| **Trigger** | External API/workflow call | Internal (manual, scheduled, or triggered) |
| **Purpose** | Single operation, quick response | Multi-step orchestration |
| **Pattern** | Stateless handler | Stateful workflow with phases |
| **Location** | `internal/worker_functions/` | `internal/jobs/` |

## Creating a Job

### Step 1: Define Job in `internal/jobs/your_job.go`

```go
package jobs

import (
	"fmt"
	"log"
	"time"
	"worker_starter_template/internal/state"
)

// YourJob orchestrates a multi-step process
type YourJob struct {
	AGS    *state.AsyncGlobalState
	Config string // Job-specific configuration
}

// YourJobResult contains execution results
type YourJobResult struct {
	ItemsProcessed int
	Success        bool
	Error          error
	ExecutionTime  time.Duration
}

// Execute runs the job
func (j *YourJob) Execute() *YourJobResult {
	result := &YourJobResult{Success: false}
	start := time.Now()

	log.Println("🚀 Starting job...")

	// Phase 1: Validate
	if err := j.validate(); err != nil {
		result.Error = err
		result.ExecutionTime = time.Since(start)
		return result
	}

	// Phase 2: Process
	processed, err := j.process()
	if err != nil {
		result.Error = err
		result.ExecutionTime = time.Since(start)
		return result
	}

	// Success
	result.ItemsProcessed = processed
	result.Success = true
	result.ExecutionTime = time.Since(start)
	log.Printf("✅ Processed %d items in %s", processed, result.ExecutionTime)

	return result
}

func (j *YourJob) validate() error {
	// Validate preconditions
	return nil
}

func (j *YourJob) process() (int, error) {
	// Your job logic here
	// Access database: j.AGS.DB.Find(&records)
	return 0, nil
}

// NewYourJob creates a job instance
func NewYourJob(ags *state.AsyncGlobalState, config string) *YourJob {
	return &YourJob{
		AGS:    ags,
		Config: config,
	}
}
```

### Step 2: Run the Job

**From worker function:**
```go
func (f *YourFunction) handler(input Input, ags *state.AsyncGlobalState) (Output, error) {
	job := jobs.NewYourJob(ags, input.Config)
	result := job.Execute()
	
	if !result.Success {
		return Output{}, result.Error
	}
	
	return Output{ItemsProcessed: result.ItemsProcessed}, nil
}
```

**From CLI or scheduled task:**
```go
func main() {
	ags, _ := state.NewAsyncGlobalState()
	job := jobs.NewYourJob(ags, "config")
	result := job.Execute()
}
```

## Job Pattern Structure

```
YourJob
├── Struct fields (config, dependencies)
├── Result struct (metrics, status)
├── Execute() - main orchestration
├── validate() - precondition checks
├── process() - core logic
├── logSuccess() - success logging
├── logError() - error logging
└── NewYourJob() - constructor
```

## Best Practices

✅ **DO:**
- Break complex work into phases
- Return detailed result structs with metrics
- Log progress at each phase
- Handle errors gracefully at each step
- Use AsyncGlobalState for database access

❌ **DON'T:**
- Mix job logic with worker function handlers
- Skip validation phases
- Lose error context (wrap with `fmt.Errorf`)
- Forget to track execution time

## Advanced: Tasks Pattern

For complex jobs, extract reusable logic into tasks:

**Task Structure (`internal/jobs/tasks/your_task.go`):**
```go
package tasks

type YourTask struct {
	AGS   *state.AsyncGlobalState
	Limit int
}

type YourTaskResult struct {
	ItemsProcessed int
	Data           []string
}

func (t *YourTask) Execute() (*YourTaskResult, error) {
	// 1. Validate
	// 2. Process
	// 3. Return results
	return &YourTaskResult{ItemsProcessed: 10}, nil
}

func NewYourTask(ags *state.AsyncGlobalState, limit int) *YourTask {
	return &YourTask{AGS: ags, Limit: limit}
}
```

**Use in Job:**
```go
func (j *YourJob) process() (int, error) {
	// Create and execute task
	task := tasks.NewYourTask(j.AGS, 100)
	result, err := task.Execute()
	if err != nil {
		return 0, err
	}
	
	// Use task results
	return result.ItemsProcessed, nil
}
```

**Benefits:**
- ✅ Reusable across multiple jobs
- ✅ Easier to test independently
- ✅ Single responsibility
- ✅ Composable (jobs can combine multiple tasks)

## Example References

- **Job:** `internal/jobs/example_job.go`
- **Task:** `internal/jobs/tasks/example_task.go`

