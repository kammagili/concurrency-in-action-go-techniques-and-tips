package main

import (
	"fmt"
	"testing"
)

func BenchmarkWorkerPoolExample(b *testing.B) {
	benchmarks := []struct {
		workers int
		tasks   int
	}{
		{workers: 1, tasks: 50_000_000},
		{workers: 10, tasks: 50_000_000},
		{workers: 50, tasks: 50_000_000},
		{workers: 100, tasks: 50_000_000},
		{workers: 200, tasks: 50_000_000},
		{workers: 500, tasks: 50_000_000},
		{workers: 1000, tasks: 50_000_000},
		{workers: 5000, tasks: 50_000_000},
	}
	for _, bm := range benchmarks {
		b.Run(
			fmt.Sprintf("%d_workers_%d_tasks", bm.workers, bm.tasks),
			func(b *testing.B) {
				for b.Loop() {
					WorkerPoolExample(bm.workers, bm.tasks)
				}
			},
		)
	}
}
