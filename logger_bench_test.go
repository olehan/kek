package kek

import (
    "github.com/olehan/kek/formatters/suggared"
    "os"
    "testing"
)

var (
    _testFactory = NewFactory(os.Stdout, suggared.NewSugaredFormatter())
    _testLogger = _testFactory.NewLogger()
    _testLoggerValues = []interface{}{"values", 2345.245245, true, false, "are so", 245245245, "values"}
)

func BenchmarkLogger_Println(b *testing.B) {
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        _testLogger.Debug.Println(_testLoggerValues...)
    }
}
