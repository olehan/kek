package sugared

import (
    "github.com/olehan/kek/ds"
    "github.com/olehan/kek/formatters"
)

func (f *SugaredTemplateFormatter) PrintTemplate(
    fs *formatters.FormatterConfig,
    template string,
    v ...interface{},
) {
    WriteMetaInfo(fs)
    f.StringifyByTemplate(fs.PoolState, template, v...)
}

func (f *SugaredTemplateFormatter) PrintTemplateMap(fs *formatters.FormatterConfig, template string, v ds.Map) {
    WriteMetaInfo(fs)
    f.StringifyByTemplateMap(fs.PoolState, template, v)
}

func (f *SugaredTemplateFormatter) PrintTemplateKeyValue(
    fs *formatters.FormatterConfig,
    template string,
    keyValues ...interface{},
) {
    WriteMetaInfo(fs)
    f.StringifyByTemplateKeyValues(fs.PoolState, template, keyValues...)
}
