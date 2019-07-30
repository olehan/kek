package kek

import (
    "github.com/olehan/kek/config"
    "github.com/olehan/kek/formatters"
)

type (
    // LoggerConfig is logger specific set of configurations
    // that also extends the default base configs.
    LoggerConfig struct {
        *config.Config
        formatter formatters.Formatter
    }
)

// NewLoggerConfig returns a new logger configuration.
func NewLoggerConfig(c *config.Config) *LoggerConfig {
    return &LoggerConfig{
        Config: c,
    }
}

// SetFormatter sets logger formatter that will be used by printers.
func (l *LoggerConfig) SetFormatter(f formatters.Formatter) *LoggerConfig {
    l.formatter = f
    return l
}

// Value returns a value of the logger config pointer.
func (l *LoggerConfig) Value() LoggerConfig {
    return *l
}

// Copy returns a new copy of the logger config pointer.
func (l *LoggerConfig) Copy() *LoggerConfig {
    cp := l.Value()
    return &cp
}
