package printer

import (
    "github.com/olehan/kek/formatters"
    "github.com/olehan/kek/pool"
    "github.com/olehan/kek/ds"
)

func (p *Printer) Print(v ...interface{}) (int, error) {
    state, fc := p.initState()
    defer p.reset(state)
    p.formatter.Print(fc, v...)
    return p.fc.Writer.Write(state.Buffer)
}

func (p *Printer) Println(v ...interface{}) (int, error) {
    state, fc := p.initState()
    defer p.reset(state)
    p.formatter.Print(fc, v...)
    state.Buffer.WriteNewLine()
    return p.fc.Writer.Write(state.Buffer)
}

func (p *Printer) PrintT(template string, v ...interface{}) (int, error) {
    state, fc := p.initState()
    defer p.reset(state)
    p.formatter.PrintTemplate(fc, template, v...)
    state.Buffer.WriteNewLine()
    return p.fc.Writer.Write(state.Buffer)
}

func (p *Printer) PrintTM(template string, v ds.Map) (int, error) {
    state, fc := p.initState()
    defer p.reset(state)
    p.formatter.PrintTemplate(fc, template, v)
    state.Buffer.WriteNewLine()
    return p.fc.Writer.Write(state.Buffer)
}

func (p *Printer) PrintTKV(template string, keyValues ...interface{}) (int, error) {
    state, fc := p.initState()
    defer p.reset(state)
    p.formatter.PrintTemplate(fc, template, keyValues...)
    state.Buffer.WriteNewLine()
    return p.fc.Writer.Write(state.Buffer)
}

func (p *Printer) getFormatterState(ps *pool.PoolState) *formatters.FormatterConfig {
    return p.fc.SetPoolState(ps).SetLevel(p.level)
}

func (p *Printer) initState() (ps *pool.PoolState, fc *formatters.FormatterConfig) {
    if p.fc.UseMutex {
        p.mutex.Lock()
    }
    ps = p.pool.Get()
    fc = p.getFormatterState(ps)
    return
}

func (p *Printer) reset(ps *pool.PoolState) {
    p.fc.SetPoolState(nil)
    p.pool.Free(ps)
    if p.fc.UseMutex {
        p.mutex.Unlock()
    }
}
