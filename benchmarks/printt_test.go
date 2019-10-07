package benchmarks

import "testing"

func BenchmarkLogger_PrintTMithEmptyWriter(b *testing.B) {
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        _testEmptyLogger.Debug.PrintT(_testSimpleFormat, _testValues...)
    }
}

func BenchmarkLogger_PrintTMithStdoutWriter(b *testing.B) {
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        _testStdoutLogger.Debug.PrintT(_testSimpleFormat, _testValues...)
    }
}

func BenchmarkLogger_PrintTMWithEmptyWriter(b *testing.B) {
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        _testEmptyLogger.Debug.PrintTM(_testMapFormat, _testMapValues)
    }
}

func BenchmarkLogger_PrintTMWithStdoutWriter(b *testing.B) {
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        _testStdoutLogger.Debug.PrintTM(_testMapFormat, _testMapValues)
    }
}
