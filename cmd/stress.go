package cmd

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type StressTestResult struct {
	duration    time.Duration
	requests    int
	statusCount map[int]int
}

func newStressTesResult(duration time.Duration, requests int, statusCount map[int]int) *StressTestResult {
	return &StressTestResult{
		duration:    duration,
		requests:    requests,
		statusCount: statusCount,
	}
}

func RunStressTest(url string, requests, concurrency int) *StressTestResult {
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, concurrency)
	statusCount := make(map[int]int)
	var statusMutex sync.Mutex
	start := time.Now()
	for i := 0; i < requests; i++ {
		wg.Add(1)
		semaphore <- struct{}{}
		go func(i int) {
			defer wg.Done()
			defer func() { <-semaphore }()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("Request #%d failed: %v\n", i+1, err)
				statusMutex.Lock()
				statusCount[0]++
				statusMutex.Unlock()
				return
			}
			resp.Body.Close()
			statusMutex.Lock()
			statusCount[resp.StatusCode]++
			statusMutex.Unlock()
		}(i)
	}
	wg.Wait()
	duration := time.Since(start)
	return newStressTesResult(duration, requests, statusCount)
}

func printReport(stressTestResult StressTestResult) {
	fmt.Println("\n===== FINAL REPORT =====")
	fmt.Printf("Total time: %v\n", stressTestResult.duration)
	fmt.Printf("Total requests: %d\n", stressTestResult.requests)

	ok200, ok := stressTestResult.statusCount[200]
	fmt.Printf("Success (200): %d\n", ok200)

	if len(stressTestResult.statusCount) > 1 || !ok {
		fmt.Println("Status code distribution:")
		for code, count := range stressTestResult.statusCount {
			if code == 0 {
				fmt.Printf("Request failures (network errors): %d\n", count)
			} else if code != 200 {
				fmt.Printf(" - %d: %d\n", code, count)
			}
		}
	}
}
