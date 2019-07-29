package printer

import (
    "github.com/olehan/kek/ds"
    "github.com/olehan/kek/formatters"
)

type (
    // LoggerPrinter is an interface that gives access to the
    // configuration functions of the printer.
    LoggerPrinter interface {
        FullPrinter
        SetFormatter(formatter formatters.Formatter) FullPrinter
    }

    // FullPrinter is an interface that describes printing
    // functionality of the printer.
    FullPrinter interface {
        BasePrinter
        TemplatePrinter
        StructuredPrinter
    }

    // BasePrinter is an interface that describes the base
    // printing functionality.
    BasePrinter interface {
        Print(values ...interface{}) FullPrinter
        Println(values ...interface{}) FullPrinter
    }

    // TemplatePrinter is an interface that describes
    // formatting/templating printer functionality.
    TemplatePrinter interface {
        PrintT(template string, values ...interface{}) FullPrinter
        PrintTM(template string, values ds.Map) FullPrinter
        PrintTKV(template string, keyValues ...interface{}) FullPrinter
    }

    // StructuredPrinter is an interface that describes
    // structured printer functionality.
    StructuredPrinter interface {
        PrintSKV(message string, keyValues ...interface{}) FullPrinter
    }
)
