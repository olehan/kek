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

// WriteMetaInfo writes meta info for minified formatter logs.
func WriteMetaInfo(fs *formatters.FormatterConfig) {
    sugared.WriteName(fs, " - ")
    sugared.WriteLevel(fs)
    fs.PoolState.Buffer.WriteString(": ")
    sugared.WriteLevelTabulation(fs)
}
