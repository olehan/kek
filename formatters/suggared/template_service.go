package suggared

import (
    "github.com/olehan/kek/ds"
    "github.com/olehan/kek/formatters"
)

func (f* SugaredTemplateFormatter) PrintTemplate(
    fs *formatters.FormatterConfig,
    template string,
    v ...interface{},
) {
    writeMetaInfo(fs)
    f.StringifyByTemplate(fs.PoolState, template, v...)
}

func (f* SugaredTemplateFormatter) PrintTemplateMap(fs *formatters.FormatterConfig, template string, v ds.Map) {
    writeMetaInfo(fs)
    f.StringifyByTemplateMap(fs.PoolState, template, v)
}

func (f* SugaredTemplateFormatter) PrintTemplateKeyValue(
    fs *formatters.FormatterConfig,
    template string,
    keyValues ...interface{},
) {
    writeMetaInfo(fs)
    f.StringifyByTemplateKeyValues(fs.PoolState, template, keyValues...)
}
