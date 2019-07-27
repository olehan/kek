package kek

import (
    "github.com/olehan/kek/config"
    "github.com/olehan/kek/formatters"
)

type (
    LoggerConfig struct {
        *config.Config
        formatter formatters.Formatter
    }
)

func NewLoggerConfig(c *config.Config) *LoggerConfig {
    return &LoggerConfig{
        Config: c,
    }
}

func (l *LoggerConfig) SetFormatter(f formatters.Formatter) *LoggerConfig {
    l.formatter = f
    return l
}

func (l *LoggerConfig) Value() LoggerConfig {
    return *l
}

func (l *LoggerConfig) Copy() *LoggerConfig {
    cp := l.Value()
    return &cp
}

func (l *LoggerConfig) CopyValue() LoggerConfig {
    return l.Copy().Value()
}
