package formatters

import (
    "github.com/olehan/kek/config"
    "github.com/olehan/kek/levels"
    "github.com/olehan/kek/pool"
)

type (
    // FormatterConfig stores formatter specific configurations
    // and extends a base config struct. FormatterConfig is passed
    // to all formatter functions.
    FormatterConfig struct {
        *config.Config
        Level     levels.Level
        PoolState *pool.State
    }
)

// NewFormatterConfig returns a new FormatterConfig.
func NewFormatterConfig(c *config.Config) *FormatterConfig {
    return &FormatterConfig{
        Config: c,
    }
}

// SetPoolState sets a new pool state into the formatter config.
func (f *FormatterConfig) SetPoolState(ps *pool.State) *FormatterConfig {
    f.PoolState = ps
    return f
}

// SetLevel sets a new level into the formatter config.
func (f *FormatterConfig) SetLevel(level levels.Level) *FormatterConfig {
    f.Level = level
    return f
}
