package main

import (
    "testing"
)

func BenchmarkRunSolution1(b *testing.B) {
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        RunSolution1(Inputs)
    }
}

func BenchmarkRunSolution2(b *testing.B) {
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        RunSolution2(Inputs)
    }
}
