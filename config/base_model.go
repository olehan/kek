package config

import (
    "github.com/olehan/kek/levels"
    "io"
)

type (
    BaseConfig struct {
        Name     string
        Writer   io.Writer
        Level    levels.Level
        UseMutex bool
    }
)

func NewBaseConfig() *BaseConfig {
    return &BaseConfig{}
}

func (c *BaseConfig) SetName(name string) *BaseConfig {
    c.Name = name
    return c
}

func (c *BaseConfig) SetUseMutex(useMutex bool) *BaseConfig {
    c.UseMutex = useMutex
    return c
}

func (c *BaseConfig) SetLevel(level levels.Level) *BaseConfig {
    c.Level = level
    return c
}

func (c *BaseConfig) SetWriter(writer io.Writer) *BaseConfig {
    c.Writer = writer
    return c
}

func (c *BaseConfig) Value() BaseConfig {
    return *c
}

func (c *BaseConfig) Copy() *BaseConfig {
    // hehe :DDDDD
    cp := c.Value()
    return &cp
}

func (c *BaseConfig) CopyValue() BaseConfig {
    return c.Copy().Value()
}
