package suggared

import "github.com/olehan/kek/formatters"

func (p *SugaredBaseFormatter) Print(fs *formatters.FormatterConfig, v ...interface{}) {
    writeMetaInfo(fs)
    p.StringifyValues(fs.PoolState, v...)
}
