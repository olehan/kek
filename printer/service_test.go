package printer

import (
    "github.com/olehan/kek/config"
    "github.com/olehan/kek/ds"
    "github.com/olehan/kek/formatters/sugared"
    "github.com/olehan/kek/levels"
    "os"
    "testing"
)

var (
    _testPrinter = NewPrinter(
        config.NewConfig().
            SetWriter(os.Stdout).
            SetUseMutex(true),
        sugared.Formatter,
        levels.Debug,
    )
    _testPrinterValues      = []interface{}{"values", 245345.345345, "are so", 24546, true, false, "values"}
    _testPrinterTemplate    = "template {} for {} that {},-- test {}, yeah {}"
    _testPrinterMapTemplate = "template {{key1}} for {{key2}}"
    _testPrinterMap         = ds.NewMap().Set("key1", "value1")
    _testPrinterKeyValues   = []interface{}{"key1", "value1", "key2", false, "unknown", 314134.134, 1462726246}
)

func TestPrinter_Print(t *testing.T) {
    _testPrinter.Print(_testPrinterValues...)
}

func TestPrinter_Println(t *testing.T) {
    _testPrinter.Println(_testPrinterValues...)
}

func TestPrinter_PrintT(t *testing.T) {
    _testPrinter.PrintT(_testPrinterTemplate, _testPrinterValues...)
}

func TestPrinter_PrintTM(t *testing.T) {
    _testPrinter.PrintTM(_testPrinterMapTemplate, _testPrinterMap)
}

func TestPrinter_PrintTKV(t *testing.T) {
    _testPrinter.PrintTKV(_testPrinterMapTemplate, _testPrinterKeyValues...)
}

func TestPrinter_PrintSKV(t *testing.T) {
    _testPrinter.PrintSKV(_testPrinterTemplate, _testPrinterKeyValues...)
}
