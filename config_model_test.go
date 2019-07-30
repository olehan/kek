package kek

import (
    "github.com/olehan/kek/config"
    "github.com/olehan/kek/formatters/minified"
    "testing"
)

func TestLoggerConfig_SetFormatter(t *testing.T) {
    c := NewLoggerConfig(config.NewConfig())
    c.SetFormatter(minified.NewMinifiedFormatter())
    if c.formatter == nil {
        panic("no formatter set")
    }
}

func TestLoggerConfig_Copy(t *testing.T) {
    c := NewLoggerConfig(config.NewConfig())
    cp := c.Copy()
    if c == cp {
        panic("config wasn't copied")
    }
}

func TestLoggerConfig_Value(t *testing.T) {
    c := NewLoggerConfig(config.NewConfig())
    cp := c.Value()
    if cp != c.Value() {
        panic("config wasn't copied")
    }
}
