package formatters

import (
    "github.com/olehan/kek/config"
    "github.com/olehan/kek/levels"
    "github.com/olehan/kek/pool"
    "testing"
)

func TestFormatterConfig_SetLevel(t *testing.T) {
    c := NewFormatterConfig(config.NewConfig())
    c.SetLevel(levels.Debug)
    if c.Level != levels.Debug {
        panic("formatter config level has not changed")
    }
}

func TestFormatterConfig_SetPoolState(t *testing.T) {
    c := NewFormatterConfig(config.NewConfig())
    c.SetPoolState(pool.NewPool().Get())
    if c.PoolState == nil {
        panic("formatter config pool state has not changed")
    }
}
