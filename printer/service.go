package printer

import (
    "github.com/olehan/kek/formatters"
    "github.com/olehan/kek/pool"
    "github.com/olehan/kek/sugar"
)

func (p *Printer) Print(v ...interface{}) (int, error) {
    state := p.initState()
    defer p.reset(state)
    p.formatter.Print(p.getFormatterState(state), v...)
    return p.Writer.Write(state.Buffer)
}

func (p *Printer) Println(v ...interface{}) (int, error) {
    state := p.initState()
    defer p.reset(state)
    p.formatter.Print(p.getFormatterState(state), v...)
    state.Buffer.WriteNewLine()
    return p.Writer.Write(state.Buffer)
}

func (p *Printer) PrintT(template string, v ...interface{}) (int, error) {
    state := p.initState()
    defer p.reset(state)
    p.formatter.PrintTemplate(p.getFormatterState(state), template, v...)
    state.Buffer.WriteNewLine()
    return p.Writer.Write(state.Buffer)
}

func (p *Printer) PrintTM(template string, v sugar.Map) (int, error) {
    state := p.initState()
    defer p.reset(state)
    p.formatter.PrintTemplate(p.getFormatterState(state), template, v)
    state.Buffer.WriteNewLine()
    return p.Writer.Write(state.Buffer)
}

func (p *Printer) PrintTKV(template string, keyValues ...interface{}) (int, error) {
    state := p.initState()
    defer p.reset(state)
    p.formatter.PrintTemplate(p.getFormatterState(state), template, keyValues...)
    state.Buffer.WriteNewLine()
    return p.Writer.Write(state.Buffer)
}

func (p *Printer) getFormatterState(ps *pool.PoolState) *formatters.FormatterConfig {
    p.fc.SetPoolState(ps)
    if p.UseMutex {
        p.mutex.Unlock()
    }
    return p.fc
}

func (p *Printer) initState() *pool.PoolState {
    if p.UseMutex {
        p.mutex.Lock()
    }
    return p.Pool.Get()
}

func (p *Printer) reset(ps *pool.PoolState) {
    p.fc.SetPoolState(nil)
    p.Pool.Free(ps)
}
