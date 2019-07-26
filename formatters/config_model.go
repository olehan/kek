package formatters

import (
    "github.com/olehan/kek/config"
    "github.com/olehan/kek/pool"
)

type (
    FormatterConfig struct {
        *config.Config
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
