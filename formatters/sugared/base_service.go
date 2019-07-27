package sugared

import "github.com/olehan/kek/formatters"

func (p *SugaredBaseFormatter) Print(fs *formatters.FormatterConfig, v ...interface{}) {
    WriteMetaInfo(fs)
    p.StringifyValues(fs.PoolState, v...)
}
