package printer

import "testing"

var (
    _testPrinter = NewPrinter()
    _testPrinterValues = []interface{}{"values", 245345.345345, "are so", 24546, true, false, "values"}
)

func TestPrinter_Print(t *testing.T) {
    _, err := _testPrinter.Print(_testPrinterValues...)
    if err != nil {
        panic("error trying to print values")
    }
}
