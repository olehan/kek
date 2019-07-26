package config

import (
    "github.com/olehan/kek/levels"
    "io"
)

type (
    Config struct {
        Name     string
        Writer   io.Writer
        Level    levels.Level
        UseMutex bool

        WithColors bool
        WithTime   bool
        WithPID    bool
    }
)

func NewConfig() *Config {
    return &Config{}
}

func (c *Config) SetName(name string) *Config {
    c.Name = name
    return c
}

func (c *Config) SetUseMutex(useMutex bool) *Config {
    c.UseMutex = useMutex
    return c
}

func (c *Config) SetLevel(level levels.Level) *Config {
    c.Level = level
    return c
}

func (c *Config) SetWriter(writer io.Writer) *Config {
    c.Writer = writer
    return c
}

func (c *Config) Value() Config {
    return *c
}

func (c *Config) Copy() *Config {
    // hehe :DDDDD
    cp := c.Value()
    return &cp
}

func (c *Config) CopyValue() Config {
    return c.Copy().Value()
}
