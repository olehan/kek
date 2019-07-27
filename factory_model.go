package kek

import (
    "github.com/olehan/kek/config"
    "github.com/olehan/kek/formatters"
    "github.com/olehan/kek/formatters/suggared"
    "io"
    "os"
    "strings"
)

type (
    Factory struct {
        *LoggerConfig
        namePrefix string
    }
)

const (
    nameSeparator = "."
    factoryLoggerNameSeparator = "/"
)

var (
    defaultFactory = NewFactory(os.Stdout, suggared.Formatter)
    defaultLogger  = defaultFactory
)

func NewFactory(writer io.Writer, formatter formatters.Formatter, name ...string) *Factory {
    return &Factory{
        namePrefix: strings.Join(name, nameSeparator),
        LoggerConfig: NewLoggerConfig(config.NewConfig().SetWriter(writer)).SetFormatter(formatter),
    }
}
