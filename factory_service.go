package kek

import (
    "github.com/olehan/kek/config"
    "github.com/olehan/kek/formatters"
    "github.com/olehan/kek/levels"
    "github.com/olehan/kek/names"
    "github.com/olehan/kek/printer"
    "strings"
)

func (f *Factory) NewLogger(name ...string) *Logger {
    n := f.registerLoggerName(name...)
    return f.newLoggerFromLeveledPrinter(f.Config.Copy().SetName(n))
}

func (f *Factory) NewLinkedLogger(name ...string) *Logger {
    n := f.registerLoggerName(name...)
    return f.newLoggerFromLeveledPrinter(f.Config.SetName(n))
}

func (f *Factory) registerLoggerName(name ...string) (n string) {
    n = strings.Join(name, nameSeparator)

    if len(f.namePrefix) > 0 {
        n = strings.Join([]string{f.namePrefix, n}, factoryLoggerNameSeparator)
    }

    names.SetLongestNameLen(len(n))
    return
}

func (f *Factory) newLoggerFromLeveledPrinter(c *config.Config) *Logger {
    thunk := f.newLeveledPrinter(c, f.formatter)
    return &Logger{
        LoggerConfig: NewLoggerConfig(c),
        Debug: thunk(levels.Debug),
        Info:  thunk(levels.Info),
        Succ:  thunk(levels.Succ),
        Note:  thunk(levels.Note),
        Error: thunk(levels.Error),
        Panic: thunk(levels.Panic),
        Warn:  thunk(levels.Warn),
        Fatal:  thunk(levels.Fatal),
    }
}

func (f *Factory) newLeveledPrinter(
    c *config.Config,
    formatter formatters.Formatter,
) func(level levels.Level) printer.PrinterRepo {
    return func(level levels.Level) printer.PrinterRepo {
        return printer.NewPrinter(c, formatter, level)
    }
}
