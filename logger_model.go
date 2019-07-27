package kek

import (
    "github.com/olehan/kek/printer"
)

type (
    Logger struct {
        *LoggerConfig

        Info    *printer.Printer
        Succ    *printer.Printer
        Debug   *printer.Printer
        Note    *printer.Printer
        Warn    *printer.Printer
        Error   *printer.Printer
        Fatal   *printer.Printer
        Panic   *printer.Printer
    }
)

func NewLogger(name ...string) *Logger {
    return defaultFactory.NewLogger(name...)
}
