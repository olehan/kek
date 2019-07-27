package printer

import (
    "github.com/olehan/kek/ds"
)

type (
    BasePrinter interface {
        Print(v ...interface{}) (int, error)
        Println(v ...interface{}) (int, error)
    }

    TemplatePrinter interface {
        PrintT(template string, v ...interface{}) (int, error)
        PrintTM(template string, v ds.Map) (int, error)
        PrintTKV(template string, keyValues ...interface{}) (int, error)
    }

    StructuredPrinter interface {
        PrintSM(v ds.Map) (int, error)
        PrintSKV(keyValues ...interface{}) (int, error)
    }
)
