package formatters

import (
    "github.com/olehan/kek/config"
    "github.com/olehan/kek/levels"
    "github.com/olehan/kek/pool"
)

type (
    FormatterConfig struct {
        *config.Config
        Level     levels.Level
        PoolState *pool.PoolState
    }
)

func NewFormatterConfig(c *config.Config) *FormatterConfig {
    return &FormatterConfig{
        Config: c,
    }
}

func (f *FormatterConfig) SetPoolState(ps *pool.PoolState) *FormatterConfig {
    f.PoolState = ps
    return f
}

func (f *FormatterConfig) SetLevel(level levels.Level) *FormatterConfig {
    f.Level = level
    return f
}
