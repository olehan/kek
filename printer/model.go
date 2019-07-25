package printer

import (
    "github.com/olehan/kek/config"
    "github.com/olehan/kek/formatters"
    "github.com/olehan/kek/formatters/suggared"
    "github.com/olehan/kek/pool"
    "os"
    "sync"
)

type (
    Printer struct {
        BasePrinter
        TemplatePrinter
        StructuredPrinter

        *config.PoolConfig
        *config.BaseConfig
        *config.SugarConfig

        formatter formatters.Formatter
        mutex     *sync.Mutex
        fc        *formatters.FormatterConfig
    }
)

func NewPrinter() *Printer {
    pc := config.NewPoolConfig().SetPool(pool.NewPool())
    bc := config.NewBaseConfig().SetWriter(os.Stdout)
    sc := config.NewSugarConfig()
    return &Printer{
        PoolConfig:  pc,
        BaseConfig:  bc,
        SugarConfig: sc,
        formatter:   suggared.NewSugaredFormatter(),
        fc:          formatters.NewFormatterConfig(pc, bc, sc),
        mutex:       &sync.Mutex{},
    }
}
