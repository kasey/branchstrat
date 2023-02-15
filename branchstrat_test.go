package main

import "testing"

var iterations = 1000000

func BenchmarkIfelse(b *testing.B) {
    for i := 0; i < b.N; i++ {
        if i == b.N / 2 {
            pivot = true
        }
        for z := 0; z < iterations; z++ {
            ifelse(z)
        }
    }
}

func BenchmarkFuncpoint(b *testing.B) {
    for i := 0; i < b.N; i++ {
        if i == b.N / 2 {
            divNum = div100
        }
        for z := 0; z < iterations; z++ {
            funcpoint(z)
        }
    }
}
