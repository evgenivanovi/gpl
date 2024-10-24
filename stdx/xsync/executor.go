package xsync

import (
	"context"
	"sync"
	"time"
)

func ExecuteSequential(tasks ...func()) {
	for _, task := range tasks {
		Run(task)
	}
}

func ExecuteParallel(tasks ...func()) {
	var wg sync.WaitGroup
	for _, task := range tasks {
		GOWG(task, &wg)
	}
	wg.Wait()
}

func ExecuteScheduled(ctx context.Context, interval time.Duration, task func()) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			{
				return
			}
		case <-ticker.C:
			{
				Run(task)
			}
		}
	}
}
