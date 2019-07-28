package printer

import (
    "github.com/olehan/kek/config"
    "github.com/olehan/kek/formatters/sugared"
    "github.com/olehan/kek/levels"
    "os"
    "testing"
)

var (
    _testPrinter = NewPrinter(config.NewConfig().SetWriter(os.Stdout), sugared.Formatter, levels.Debug)
    _testPrinterValues = []interface{}{"values", 245345.345345, "are so", 24546, true, false, "values"}
)

func TestPrinter_Print(t *testing.T) {
    _testPrinter.Print(_testPrinterValues...)
}
