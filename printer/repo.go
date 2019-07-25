package printer

import (
    "github.com/olehan/kek/sugar"
)

type (
    BasePrinter interface {
        Print(v ...interface{}) (int, error)
        Println(v ...interface{}) (int, error)
    }

    TemplatePrinter interface {
        PrintT(template string, v ...interface{}) (int, error)
        PrintTM(template string, v sugar.Map) (int, error)
        PrintTKV(template string, keyValues ...interface{}) (int, error)
    }

    StructuredPrinter interface {
        PrintSM(v sugar.Map) (int, error)
        PrintSKV(keyValues ...interface{}) (int, error)
    }
)
