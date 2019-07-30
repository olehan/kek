package config

import (
    "github.com/olehan/kek/colors"
    "os"
    "testing"
)

func TestConfig_Copy(t *testing.T) {
    c := NewConfig()
    cp := c.Copy()
    if c == cp {
        panic("config wasn't copied")
    }
}

func TestConfig_Value(t *testing.T) {
    c := NewConfig()
    cp := c.Value()
    if cp != c.Value() {
        panic("config wasn't copied")
    }
}

func TestConfig_SetName(t *testing.T) {
    expectedValue := "name"
    c := NewConfig()
    c.SetName(expectedValue)
    if expectedValue != c.Name {
        panic("config couldn't set the value")
    }
}

func TestConfig_SetNameColor(t *testing.T) {
    clrs := []colors.Color{colors.Black, colors.Bold}
    expectedValue := colors.String(clrs...)
    c := NewConfig()
    c.SetNameColor(clrs...)
    if expectedValue != c.NameColor {
        panic("config couldn't set the value")
    }
}

func TestConfig_SetRandomNameColor(t *testing.T) {
    c := NewConfig()
    c.SetRandomNameColor()
    if len(c.NameColor) == 0 {
        panic("config couldn't generate a random color")
    }
}

func TestConfig_SetUseMutex(t *testing.T) {
    expectedValue := !defaultConfig.UseMutex
    c := NewConfig()
    c.SetUseMutex(expectedValue)
    if expectedValue != c.UseMutex {
        panic("config couldn't set the value")
    }
}

func TestConfig_SetWithColors(t *testing.T) {
    expectedValue := !defaultConfig.WithColors
    c := NewConfig()
    c.SetWithColors(expectedValue)
    if expectedValue != c.WithColors {
        panic("config couldn't set the value")
    }
}

func TestConfig_SetWithDateTime(t *testing.T) {
    expectedDate := !defaultConfig.WithDate
    expectedTime := !defaultConfig.WithTime
    expectedNS := !defaultConfig.WithNS
    c := NewConfig()
    c.SetWithDateTime(expectedDate, expectedTime, expectedNS)
    if expectedDate != c.WithDate || expectedTime != c.WithTime || expectedNS != c.WithNS {
        panic("config couldn't set the value")
    }
}

func TestConfig_SetWithNameTabulation(t *testing.T) {
    expectedValue := !defaultConfig.WithNameTabulation
    c := NewConfig()
    c.SetWithNameTabulation(expectedValue)
    if expectedValue != c.WithNameTabulation {
        panic("config couldn't set the value")
    }
}

func TestConfig_SetWithPID(t *testing.T) {
    expectedValue := !defaultConfig.WithPID
    c := NewConfig()
    c.SetWithPID(expectedValue)
    if expectedValue != c.WithPID {
        panic("config couldn't set the value")
    }
}

func TestConfig_SetWriter(t *testing.T) {
    c := NewConfig()
    c.SetWriter(os.Stderr)
    if os.Stderr != c.Writer {
        panic("config couldn't set the value")
    }
}
