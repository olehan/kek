package printer

import (
    "github.com/olehan/kek/config"
    "github.com/olehan/kek/formatters"
    "github.com/olehan/kek/formatters/suggared"
    "github.com/olehan/kek/pool"
    "sync"
)

type (
    Printer struct {
        formatter formatters.Formatter
        mutex     *sync.Mutex
        pool      *pool.Pool
        fc        *formatters.FormatterConfig
    }
)

func NewPrinter(c *config.Config) *Printer {
    return &Printer{
        mutex:     &sync.Mutex{},
        pool:      pool.NewPool(),
        formatter: suggared.NewSugaredFormatter(),
        fc:        formatters.NewFormatterConfig(c),
    }
}

func (p *Printer) SetFormatter(formatter formatters.Formatter) *Printer {
    p.formatter = formatter
    return p
}
