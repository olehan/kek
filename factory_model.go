package kek

import (
    "github.com/olehan/kek/config"
    "github.com/olehan/kek/formatters"
    "github.com/olehan/kek/formatters/sugared"
    "io"
    "os"
    "strings"
)

type (
    // Factory is an entity that holds the logger configurations
    // and populates Loggers with an option of linking the config
    // or only extending it.
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
    defaultFactory = NewFactory(os.Stdout, sugared.Formatter)
)

// NewFactory returns a new factory with a set of required configurations like
// writer and formatter.
func NewFactory(writer io.Writer, formatter formatters.Formatter, name ...string) *Factory {
    return &Factory{
        namePrefix: strings.Join(name, nameSeparator),
        LoggerConfig: NewLoggerConfig(config.NewConfig().SetWriter(writer)).SetFormatter(formatter),
    }
}
