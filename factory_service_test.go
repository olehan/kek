package kek

import (
    "github.com/olehan/kek/formatters/minified"
    "os"
    "testing"
)

func TestFactory_NewLogger(t *testing.T) {
    f := NewFactory(os.Stdout, minified.Formatter, "name")
    expectedValue := false
    f.SetUseMutex(!expectedValue)

    l := f.NewLogger()
    if l.UseMutex == expectedValue {
        panic("logger didn't extend factory config")
    }

    f.SetUseMutex(expectedValue)
    if l.UseMutex == expectedValue {
        panic("logger is linked to the factory config")
    }
}

func TestFactory_NewLinkedLogger(t *testing.T) {
    f := NewFactory(os.Stdout, minified.Formatter, "name")
    expectedValue := false
    f.SetUseMutex(!expectedValue)

    l := f.NewLinkedLogger()
    if l.UseMutex == expectedValue {
        panic("logger didn't extend factory config")
    }

    f.SetUseMutex(expectedValue)
    if l.UseMutex != expectedValue {
        panic("logger didn't linked to the factory config")
    }
}
