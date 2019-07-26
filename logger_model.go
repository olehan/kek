package kek

import (
    "github.com/olehan/kek/config"
    "github.com/olehan/kek/printer"
)

type (
    Logger struct {
        *config.Config

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

func NewLogger() *Logger {
    return defaultFactory.NewLogger()
}
