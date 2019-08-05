package printer

import (
    "github.com/olehan/kek/ds"
    "github.com/olehan/kek/formatters"
    "github.com/olehan/kek/levels"
    "github.com/olehan/kek/pool"
    "os"
)

// Print manages base formatters writing functionality.
func (p *Printer) Print(v ...interface{}) FullPrinter {
    state, fc := p.initState()
    defer p.reset(state)
    p.formatter.Print(fc, v...)
    p.fc.Writer.Write(state.Buffer)
    return p
}

// Println manages base formatters writing functionality
// adding a new line to the end of the buffer.
func (p *Printer) Println(v ...interface{}) FullPrinter {
    state, fc := p.initState()
    defer p.reset(state)
    p.formatter.Print(fc, v...)
    state.Buffer.WriteNewLine()
    p.fc.Writer.Write(state.Buffer)
    return p
}

// PrintT manages template formatters writing functionality
// adding a new line to the end of the buffer.
func (p *Printer) PrintT(template string, v ...interface{}) FullPrinter {
    state, fc := p.initState()
    defer p.reset(state)
    p.formatter.PrintTemplate(fc, template, v...)
    state.Buffer.WriteNewLine()
    p.fc.Writer.Write(state.Buffer)
    return p
}

// PrintTM manages map template formatters writing functionality
// adding a new line to the end of the buffer.
func (p *Printer) PrintTM(template string, v ds.Map) FullPrinter {
    state, fc := p.initState()
    defer p.reset(state)
    p.formatter.PrintTemplateMap(fc, template, v)
    state.Buffer.WriteNewLine()
    p.fc.Writer.Write(state.Buffer)
    return p
}

// PrintTKV manages key value template formatters writing functionality
// adding a new line to the end of the buffer.
func (p *Printer) PrintTKV(template string, keyValues ...interface{}) FullPrinter {
    state, fc := p.initState()
    defer p.reset(state)
    p.formatter.PrintTemplateKeyValue(fc, template, keyValues...)
    state.Buffer.WriteNewLine()
    p.fc.Writer.Write(state.Buffer)
    return p
}

// PrintSKV manages structured key value formatters writing functionality
// adding a new line to the end of the buffer.
func (p *Printer) PrintSKV(message string, keyValues ...interface{}) FullPrinter {
    state, fc := p.initState()
    defer p.reset(state)
    p.formatter.PrintStructKeyValues(fc, message, keyValues...)
    p.fc.Writer.Write(state.Buffer)
    return p
}

func (p *Printer) getFormatterState(ps *pool.State) *formatters.FormatterConfig {
    return p.fc.SetPoolState(ps).SetLevel(p.level)
}

func (p *Printer) initState() (ps *pool.State, fc *formatters.FormatterConfig) {
    if p.fc.UseMutex {
        p.mutex.Lock()
    }
    ps = p.pool.Get()
    fc = p.getFormatterState(ps)
    return
}

func (p *Printer) reset(ps *pool.State) {
    switch p.level {
    case levels.Fatal:
        os.Exit(1)
    case levels.Panic:
        panic(string(ps.Buffer))
    default:
        p.fc.SetPoolState(nil)
        p.pool.Free(ps)
        if p.fc.UseMutex {
            p.mutex.Unlock()
        }
    }
}
