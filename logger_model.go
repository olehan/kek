package kek

import (
    "github.com/olehan/kek/formatters"
    "github.com/olehan/kek/printer"
)

type (
    Logger struct {
        *LoggerConfig

        Info    printer.PrinterRepo
        Succ    printer.PrinterRepo
        Debug   printer.PrinterRepo
        Note    printer.PrinterRepo
        Warn    printer.PrinterRepo
        Error   printer.PrinterRepo
        Fatal   printer.PrinterRepo
        Panic   printer.PrinterRepo
    }
)

func NewLogger(name ...string) *Logger {
    return defaultFactory.NewLogger(name...)
}

func (l *Logger) SetFormatter(formatter formatters.Formatter) *Logger {
    l.Debug.(printer.LoggerPrinterRepo).SetFormatter(formatter)
    l.Info.(printer.LoggerPrinterRepo).SetFormatter(formatter)
    l.Warn.(printer.LoggerPrinterRepo).SetFormatter(formatter)
    l.Fatal.(printer.LoggerPrinterRepo).SetFormatter(formatter)
    l.Error.(printer.LoggerPrinterRepo).SetFormatter(formatter)
    l.Panic.(printer.LoggerPrinterRepo).SetFormatter(formatter)
    l.Note.(printer.LoggerPrinterRepo).SetFormatter(formatter)
    l.Succ.(printer.LoggerPrinterRepo).SetFormatter(formatter)
    return l
}
