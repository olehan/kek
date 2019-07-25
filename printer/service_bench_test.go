package printer

import "testing"

func BenchmarkPrinter_Println(b *testing.B) {
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        _testPrinter.Println(_testPrinterValues...)
    }
}
