package main

import (
	"math/rand"
	"testing"
)

func Benchmark(b *testing.B) {
	memTable := New(41000)

	for i := 0; i < b.N; i++ {
		memTable.Append(rand.Intn(10000), rand.Intn(10000))
	}
}
