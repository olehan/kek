package kek

import (
    "github.com/olehan/kek/config"
    "github.com/olehan/kek/levels"
    "github.com/olehan/kek/printer"
)

func (f *Factory) NewLogger() *Logger {
    return f.newLoggerFromLeveledPrinter(f.Config.Copy())
}

func (f *Factory) NewLinkedLogger() *Logger {
    return f.newLoggerFromLeveledPrinter(f.Config)
}

func (f *Factory) newLoggerFromLeveledPrinter(c *config.Config) *Logger {
    thunk := f.newLeveledPrinter(c)
    return &Logger{
        Config: c,
        Debug: thunk(levels.Debug),
        Info:  thunk(levels.Info),
        Succ:  thunk(levels.Succ),
        Note:  thunk(levels.Note),
        Error: thunk(levels.Error),
        Panic: thunk(levels.Panic),
        Warn:  thunk(levels.Warn),
    }
}

func (f *Factory) newLeveledPrinter(c *config.Config) func(level levels.Level) *printer.Printer {
    return func(level levels.Level) *printer.Printer {
        return printer.NewPrinter(c.SetLevel(level))
    }
}
