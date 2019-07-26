package kek

import (
    "github.com/olehan/kek/config"
    "github.com/olehan/kek/formatters"
    "github.com/olehan/kek/formatters/suggared"
    "io"
    "os"
)

type (
    Factory struct {
        *config.Config
        formatter formatters.Formatter
    }
)

var (
    defaultFactory = NewFactory(os.Stdout, suggared.NewSugaredFormatter())
)

func NewFactory(w io.Writer, f formatters.Formatter) *Factory {
    return &Factory{
        formatter: f,
        Config:    config.NewConfig().SetWriter(w),
    }
}
