package kek

import (
    "github.com/olehan/kek/formatters"
    "github.com/olehan/kek/printer"
)

type (
    // Logger is an entity that configures and holds leveled printers.
    Logger struct {
        *LoggerConfig

        Info    printer.FullPrinter
        Succ    printer.FullPrinter
        Debug   printer.FullPrinter
        Note    printer.FullPrinter
        Warn    printer.FullPrinter
        Error   printer.FullPrinter
        Fatal   printer.FullPrinter
        Panic   printer.FullPrinter
    }
)

// NewLogger returns a new logger created from a default factory.
func NewLogger(name ...string) *Logger {
    return defaultFactory.NewLogger(name...)
}

// SetFormatter sets one formatter across the all leveled printers.
func (l *Logger) SetFormatter(formatter formatters.Formatter) *Logger {
    l.Debug.(printer.LoggerPrinter).SetFormatter(formatter)
    l.Info.(printer.LoggerPrinter).SetFormatter(formatter)
    l.Warn.(printer.LoggerPrinter).SetFormatter(formatter)
    l.Fatal.(printer.LoggerPrinter).SetFormatter(formatter)
    l.Error.(printer.LoggerPrinter).SetFormatter(formatter)
    l.Panic.(printer.LoggerPrinter).SetFormatter(formatter)
    l.Note.(printer.LoggerPrinter).SetFormatter(formatter)
    l.Succ.(printer.LoggerPrinter).SetFormatter(formatter)
    l.formatter = formatter
    return l
}
