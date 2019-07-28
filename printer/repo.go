package printer

import (
    "github.com/olehan/kek/ds"
    "github.com/olehan/kek/formatters"
)

type (
    LoggerPrinterRepo interface {
        PrinterRepo
        SetFormatter(formatter formatters.Formatter) PrinterRepo
    }

    PrinterRepo interface {
        BasePrinter
        TemplatePrinter
        StructuredPrinter
    }

    BasePrinter interface {
        Print(values ...interface{}) PrinterRepo
        Println(values ...interface{}) PrinterRepo
    }

    TemplatePrinter interface {
        PrintT(template string, values ...interface{}) PrinterRepo
        PrintTM(template string, values ds.Map) PrinterRepo
        PrintTKV(template string, keyValues ...interface{}) PrinterRepo
    }

    StructuredPrinter interface {
        PrintSKV(message string, keyValues ...interface{}) PrinterRepo
    }
)
