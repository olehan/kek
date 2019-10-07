package benchmarks

import (
    "testing"
)

func BenchmarkLogger_PrintlnWithEmptyWriter(b *testing.B) {
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        _testEmptyLogger.Debug.Println(_testSimpleMock...)
    }
}

func BenchmarkLogger_PrintlnWithStdoutWriter(b *testing.B) {
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        _testStdoutLogger.Debug.Println(_testSimpleMock...)
    }
}
