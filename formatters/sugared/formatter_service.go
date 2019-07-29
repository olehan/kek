package sugared

import (
    "github.com/olehan/kek/ds"
    "github.com/olehan/kek/formatters"
)

func (s *FormatterModel) Print(fs *formatters.FormatterConfig, v ...interface{}) {
    WriteMetaInfo(fs)
    s.StringifyValues(fs.PoolState, v...)
}

func (s *FormatterModel) PrintTemplate(
    fs *formatters.FormatterConfig,
    template string,
    v ...interface{},
) {
    WriteMetaInfo(fs)
    s.StringifyByTemplate(fs.PoolState, template, v...)
}

func (s *FormatterModel) PrintTemplateMap(
    fs *formatters.FormatterConfig,
    template string,
    v ds.Map,
) {
    WriteMetaInfo(fs)
    s.StringifyByTemplateMap(fs.PoolState, template, v)
}

func (s *FormatterModel) PrintTemplateKeyValue(
    fs *formatters.FormatterConfig,
    template string,
    keyValues ...interface{},
) {
    s.KeyValuesToMap(fs.PoolState, keyValues...)
    WriteMetaInfo(fs)
    s.StringifyByTemplateMap(fs.PoolState, template, fs.PoolState.Map)
}

func (s *FormatterModel) PrintStructKeyValues(
    fs *formatters.FormatterConfig,
    message string,
    keyValues ...interface{},
) {
    WriteMetaInfo(fs)

    fs.PoolState.Buffer.WriteString(message)
    fs.PoolState.Buffer.WriteNewLine()

    keyValuesLen := len(keyValues)

    for i := 0; i < keyValuesLen; i += 2 {
        if keyValuesLen < i + 2 {
            break
        }

        key := keyValues[i]

        if fs.WithColors {
            fs.PoolState.Buffer.WriteString(bold)
        }

        fs.PoolState.Buffer.WriteString(" > ")
        switch key.(type) {
        case string:
            fs.PoolState.Buffer.WriteString(key.(string))
        default:
            s.StringifyValue(fs.PoolState, key)
        }
        fs.PoolState.Buffer.WriteString(": ")

        if fs.WithColors {
            fs.PoolState.Buffer.WriteString(reset)
        }

        s.StringifyValue(fs.PoolState, keyValues[i + 1])
        fs.PoolState.Buffer.WriteNewLine()
    }
}
