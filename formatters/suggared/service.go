package suggared

import (
    "github.com/olehan/kek/formatters"
)

func writeMetaInfo(fs *formatters.FormatterConfig) {
    if len(fs.Name) > 0 {
        fs.PoolState.Buffer.WriteString(fs.Name)
        fs.PoolState.Buffer.WriteString(" | ")
    }
}
