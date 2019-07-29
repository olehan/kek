package printer

import (
    "github.com/olehan/kek/config"
    "github.com/olehan/kek/formatters"
    "github.com/olehan/kek/levels"
    "github.com/olehan/kek/pool"
    "sync"
)

type (
    // Printer is an entity that controls formatters printing functionality
    // and manages pool creation/reset
    Printer struct {
        level     levels.Level
        formatter formatters.Formatter
        mutex     *sync.Mutex
        pool      *pool.Pool
        fc        *formatters.FormatterConfig
    }
)

// NewPrinter returns a new printer repo that separates printing functions
// with config setters to access it letter in logger.
func NewPrinter(c *config.Config, f formatters.Formatter, level levels.Level) FullPrinter {
    return &Printer{
        formatter: f,
        level:     level,
        mutex:     &sync.Mutex{},
        pool:      pool.NewPool(),
        fc:        formatters.NewFormatterConfig(c),
    }
}

// SetFormatter sets a new formatter into the printer.
func (p *Printer) SetFormatter(formatter formatters.Formatter) FullPrinter {
    p.formatter = formatter
    return p
}
