package kek

import (
    "github.com/olehan/kek/ds"
    "github.com/olehan/kek/formatters/sugared"
    "os"
    "testing"
)

var (
    _testFactory = NewFactory(os.Stdout, sugared.Formatter)
    _testLogger = _testFactory.NewLogger("name")
    _testLoggerValues = []interface{}{"values", 2345.245245, true, false, "are so", 245245245, "values"}
)

func BenchmarkLogger_Println(b *testing.B) {
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        _testLogger.Debug.Println(_testLoggerValues...)
    }
}

func BenchmarkLogger_PrintT(b *testing.B) {
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        _testLogger.Debug.PrintT("val {} format {} template- {} asasd {} - {}sdasd {},,, {}", _testLoggerValues...)
    }
}

func BenchmarkLogger_PrintTM(b *testing.B) {
    values := ds.Map{
        "key1": "asdafgagwrt",
        "key2": 34635636.346356,
        "key3": 34234624624,
        "key4": "asd",
    }

    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        _testLogger.Debug.PrintTM("val {{key1}} format {{key2}} template- {{key3}} asasd {{key4}}-asd", values)
    }
}

func BenchmarkLogger_PrintTKV(b *testing.B) {
    values := []interface{}{
        "key1", "value1",
        "key2", "value2",
        "key3", "value3",
        "key4", "value4",
        "unknown", 25245,
        245245,
    }

    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        _testLogger.Debug.PrintTKV("there is a {{key1}} and {{key2}}, also {{key3}} asdo-a- {{key4}}", values...)
    }
}
