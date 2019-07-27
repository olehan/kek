package minified

import (
    "github.com/olehan/kek/ds"
    "github.com/olehan/kek/formatters"
    "github.com/olehan/kek/formatters/sugared"
)

func (m *MinifiedFormatter) Print(fs *formatters.FormatterConfig, v ...interface{}) {
    WriteMetaInfo(fs)
    m.StringifyValues(fs.PoolState, v...)
}

func (m *MinifiedFormatter) PrintTemplate(
    fs *formatters.FormatterConfig,
    template string,
    v ...interface{},
) {
    WriteMetaInfo(fs)
    m.StringifyByTemplate(fs.PoolState, template, v...)
}

func (m *MinifiedFormatter) PrintTemplateMap(fs *formatters.FormatterConfig, template string, v ds.Map) {
    WriteMetaInfo(fs)
    m.StringifyByTemplateMap(fs.PoolState, template, v)
}

func (m *MinifiedFormatter) PrintTemplateKeyValue(
    fs *formatters.FormatterConfig,
    template string,
    keyValues ...interface{},
) {
    WriteMetaInfo(fs)
    m.StringifyByTemplateKeyValues(fs.PoolState, template, keyValues...)
}

func WriteMetaInfo(fs *formatters.FormatterConfig) {
    sugared.WriteName(fs, " - ")
    sugared.WriteLevel(fs)
    fs.PoolState.Buffer.WriteString(": ")
}
