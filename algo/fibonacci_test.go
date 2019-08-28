package algo

import (
	"testing"
)

func BenchmarkFibo(b *testing.B) {
	for i := 0; i < 40; i++ {
		Fibo(i)
	}
}

func BenchmarkFibo2(b *testing.B) {
	for i := 0; i < 40; i++ {
		Fibo2(i)
	}
}

func BenchmarkFibo3(b *testing.B) {
	memo := make(map[int]int)
	for i := 0; i < 40; i++ {
		Memorize(i, memo)
	}
}
