package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Job represents a task to be processed
type Job struct {
	ID   int
	Data string
}

// Result represents the output of processing a job
type Result struct {
	JobID int
	Value int
}

// worker function that processes jobs from a channel
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job.ID)

		// Simulate some work (random processing time)
		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)

		// Send result back
		results <- Result{
			JobID: job.ID,
			Value: len(job.Data) * 10, // Simple calculation
		}
	}
}

func main() {
	const numWorkers = 3
	const numJobs = 10

	// Create channels
	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	// WaitGroup to wait for all workers to finish
	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{
			ID:   j,
			Data: fmt.Sprintf("task-data-%d", j),
		}
	}
	close(jobs) // Important: close the jobs channel

	// Collect results in a separate goroutine
	go func() {
		wg.Wait()      // Wait for all workers to finish
		close(results) // Close results channel
	}()

	// Process results
	fmt.Println("\nResults:")
	for result := range results {
		fmt.Printf("Job %d completed with value: %d\n", result.JobID, result.Value)
	}

	fmt.Println("All jobs completed!")
}
