package main

import "testing"

func BenchmarkProcessData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProcessData(testData)
	}
}

func BenchmarkProcessDataWorkerPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProcessDataWorkerPool(testData)
	}
}

func BenchmarkProcessMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProcessDataMutex(testData)
	}
}
