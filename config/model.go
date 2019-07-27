package config

import (
    "github.com/olehan/kek/colors"
    "io"
    "math/rand"
)

type (
    Config struct {
        Name       string
        Writer     io.Writer
        UseMutex   bool

        WithColors bool
        WithPID    bool

        WithDate   bool
        WithTime   bool
        WithNS     bool

        NameColor  string
        WithNameTabulation bool
    }
)

var (
    defaultConfig = &Config{
        WithColors:         true,
        WithPID:            true,
        WithDate:           true,
        WithTime:           true,
        WithNS:             true,
        WithNameTabulation: true,
    }
    nameColors = []colors.Color{
        colors.Red,
        colors.Green,
        colors.Yellow,
        colors.Blue,
        colors.Magenta,
        colors.Cyan,
    }
)

func NewConfig() *Config {
    return defaultConfig.Copy()
}

func (c *Config) SetName(name string) *Config {
    c.Name = name
    return c
}

func (c *Config) SetUseMutex(useMutex bool) *Config {
    c.UseMutex = useMutex
    return c
}

func (c *Config) SetWriter(writer io.Writer) *Config {
    c.Writer = writer
    return c
}

func (c *Config) SetRandomNameColor() *Config {
    c.SetNameColor(nameColors[rand.Intn(len(nameColors))])
    return c
}

func (c *Config) SetNameColor(color ...colors.Color) *Config {
    c.NameColor = colors.String(color...)
    return c
}

func (c *Config) SetWithColors(v bool) *Config {
    c.WithColors = v
    return c
}

func (c *Config) SetWithNameTabulation(v bool) *Config {
    c.WithNameTabulation = v
    return c
}

func (c *Config) SetWithPID(v bool) *Config {
    c.WithPID = v
    return c
}

func (c *Config) SetWithDateTime(date bool, time bool, ms bool) *Config {
    return c.SetWithDate(date).SetWithTime(time).SetWithNS(ms)
}

func (c *Config) SetWithDate(date bool) *Config {
    c.WithDate = date
    return c
}

func (c *Config) SetWithTime(time bool) *Config {
    c.WithTime = time
    return c
}

func (c *Config) SetWithNS(ns bool) *Config {
    c.WithNS = ns
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
