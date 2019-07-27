package minified

import "github.com/olehan/kek/formatters"

func (m *MinifiedFormatter) Print(fs *formatters.FormatterConfig, v ...interface{}) {
    m.StringifyValues(fs.PoolState, v...)
}
