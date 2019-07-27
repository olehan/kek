package printer

import (
    "github.com/olehan/kek/config"
    "github.com/olehan/kek/formatters"
    "github.com/olehan/kek/levels"
    "github.com/olehan/kek/pool"
    "sync"
)

type (
    Printer struct {
        level     levels.Level
        formatter formatters.Formatter
        mutex     *sync.Mutex
        pool      *pool.Pool
        fc        *formatters.FormatterConfig
    }
)

func NewPrinter(c *config.Config, f formatters.Formatter, level levels.Level) PrinterRepo {
    return &Printer{
        formatter: f,
        level:     level,
        mutex:     &sync.Mutex{},
        pool:      pool.NewPool(),
        fc:        formatters.NewFormatterConfig(c),
    }
}

func (p *Printer) SetFormatter(formatter formatters.Formatter) PrinterRepo {
    p.formatter = formatter
    return p
}
