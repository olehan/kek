package minified

import (
    "github.com/olehan/kek/colors"
    "github.com/olehan/kek/formatters"
    "github.com/olehan/kek/formatters/sugared"
)

var (
    reset = colors.Reset.String()
    bold  = colors.Bold.String()
)

func WriteMetaInfo(fs *formatters.FormatterConfig) {
    sugared.WriteName(fs, " - ")
    sugared.WriteLevel(fs)
    fs.PoolState.Buffer.WriteString(": ")
}
