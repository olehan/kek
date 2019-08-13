package cli

import (
    "github.com/olehan/kek/colors"
    "github.com/olehan/kek/formatters"
    "github.com/olehan/kek/formatters/sugared"
    "time"
)

var (
    reset = colors.Reset.String()
    bold  = colors.Bold.String()
)

func (m *FormatterModel) WriteMetaInfo(fs *formatters.FormatterConfig) {
    if len(m.appName) > 0 {
        fs.PoolState.Buffer.WriteString(m.appName)
        fs.PoolState.Buffer.WriteString(": ")
    }

    sugared.WriteName(fs, " - ")

    if fs.WithColors {
        fs.PoolState.Buffer.WriteString(bold)
    }

    WriteTime(fs)

    if fs.WithColors {
        fs.PoolState.Buffer.WriteString(reset)
    }

    sugared.WriteLevel(fs)
    sugared.WriteLevelTabulation(fs)
    fs.PoolState.Buffer.WriteString(" | ")
}

func WriteTime(fs *formatters.FormatterConfig) {
    now := time.Now()

    if fs.WithTime {
        fs.PoolState.Buffer.WriteInt(int64(now.Hour()))
        fs.PoolState.Buffer.WriteByte(':')
        fs.PoolState.Buffer.WriteInt(int64(now.Minute()))
        fs.PoolState.Buffer.WriteByte(':')
        fs.PoolState.Buffer.WriteInt(int64(now.Second()))
    }

    fs.PoolState.Buffer.WriteSpace()
}
