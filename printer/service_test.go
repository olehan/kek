package printer

import (
    "github.com/olehan/kek/config"
    "os"
    "testing"
)

var (
    _testPrinter = NewPrinter(config.NewConfig().SetWriter(os.Stdout))
    _testPrinterValues = []interface{}{"values", 245345.345345, "are so", 24546, true, false, "values"}
)

func TestPrinter_Print(t *testing.T) {
    _, err := _testPrinter.Print(_testPrinterValues...)
    if err != nil {
        panic("error trying to print values")
    }
}
