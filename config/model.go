package config

import (
    "github.com/olehan/kek/colors"
    "github.com/olehan/kek/names"
    "io"
    "math/rand"
)

type (
    // Config is an entity that stores the base configurations
    // with a basic setter function for every property.
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

// NewConfig returns a copy of a default Config.
func NewConfig() *Config {
    return defaultConfig.Copy()
}

// SetName sets a new name into the config
// that will be used by most formatters.
func (c *Config) SetName(name string) *Config {
    c.Name = name
    names.SetLongestNameLen(len(name))
    return c
}

// SetUseMutex sets a new value for UseMutex property.
func (c *Config) SetUseMutex(useMutex bool) *Config {
    c.UseMutex = useMutex
    return c
}

// SetWriter sets a new writer that will be used by printers.
func (c *Config) SetWriter(writer io.Writer) *Config {
    c.Writer = writer
    return c
}

// SetRandomNameColor generates a non-negative pseudo-random int
// between [0, len(nameColors)] to decide which color is going
// to be used in most formatters that support colors.
func (c *Config) SetRandomNameColor() *Config {
    c.SetNameColor(nameColors[rand.Intn(len(nameColors))])
    return c
}

// SetNameColor joins and stringify's colors and sets the final
// string into a config.
func (c *Config) SetNameColor(color ...colors.Color) *Config {
    c.NameColor = colors.String(color...)
    return c
}
// SetWithColors sets a new value for WithColors property.
func (c *Config) SetWithColors(v bool) *Config {
    c.WithColors = v
    return c
}

// SetWithNameTabulation sets a new value for WithNameTabulation property.
func (c *Config) SetWithNameTabulation(v bool) *Config {
    c.WithNameTabulation = v
    return c
}

// SetWithPID sets a new value for WithPID property.
func (c *Config) SetWithPID(v bool) *Config {
    c.WithPID = v
    return c
}

// SetWithDateTime is a shortcut for SetWithDate, SetWithTime and SetWithNS.
func (c *Config) SetWithDateTime(date bool, time bool, ms bool) *Config {
    return c.SetWithDate(date).SetWithTime(time).SetWithNS(ms)
}

// SetWithDate sets a new value for WithDate property.
func (c *Config) SetWithDate(date bool) *Config {
    c.WithDate = date
    return c
}

// SetWithTime sets a new value for WithTime property.
func (c *Config) SetWithTime(time bool) *Config {
    c.WithTime = time
    return c
}

// SetWithNS sets a new value for WithNS property.
func (c *Config) SetWithNS(ns bool) *Config {
    c.WithNS = ns
    return c
}

// Value returns a value of the config pointer.
func (c *Config) Value() Config {
    return *c
}

// Copy returns a new copy of the config pointer.
func (c *Config) Copy() *Config {
    // hehe :DDDDD
    cp := c.Value()
    return &cp
}

// CopyValue returns a new copy of the config value.
func (c *Config) CopyValue() Config {
    return c.Copy().Value()
}
