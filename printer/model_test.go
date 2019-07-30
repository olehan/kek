package printer

import (
    "github.com/olehan/kek/formatters/minified"
    "testing"
)

func TestPrinter_SetFormatter(t *testing.T) {
    loggerPrinter := _testPrinter.(LoggerPrinter)
    loggerPrinter.SetFormatter(minified.NewMinifiedFormatter())
}
