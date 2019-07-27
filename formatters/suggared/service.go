package suggared

import (
    "github.com/olehan/kek/colors"
    "github.com/olehan/kek/formatters"
    "github.com/olehan/kek/levels"
    "github.com/olehan/kek/names"
    "os"
    "time"
)

var (
    reset = colors.Reset.String()
    dateTimeColor = colors.Bold.String()
)

func writeMetaInfo(fs *formatters.FormatterConfig) {
    writeName(fs)
    writeLevel(fs)

    if fs.WithColors {
        fs.PoolState.Buffer.WriteString(dateTimeColor)
    }

    fs.PoolState.Buffer.WriteSpace()

    writeDateTime(fs)

    if fs.WithColors {
        fs.PoolState.Buffer.WriteString(reset)
    }

    writePID(fs)
    fs.PoolState.Buffer.WriteString(":   ")
}

func writeName(fs *formatters.FormatterConfig)  {
    if len(fs.Name) > 0 {
        shouldWriteColor := fs.WithColors && len(fs.NameColor) > 0

        if shouldWriteColor {
            fs.PoolState.Buffer.WriteString(fs.NameColor)
        }

        fs.PoolState.Buffer.WriteString(fs.Name)

        if fs.WithNameTabulation {
            for i := 0; i < names.LongestNameLen() - len(fs.Name); i++ {
                fs.PoolState.Buffer.WriteSpace()
            }
        }

        fs.PoolState.Buffer.WriteString(" | ")

        if shouldWriteColor {
            fs.PoolState.Buffer.WriteString(reset)
        }
    }
}

func writeLevel(fs *formatters.FormatterConfig)  {
    if fs.WithColors {
        fs.PoolState.Buffer.WriteString(levels.ColoredLevelMap[fs.Level])
    } else {
        fs.PoolState.Buffer.WriteString(levels.NonColoredLevelMap[fs.Level])
    }
}

func writeDateTime(fs *formatters.FormatterConfig)  {
    now := time.Now()

    if fs.WithDate {
        year, month, day := now.Date()
        fs.PoolState.Buffer.WriteInt(int64(year))
        fs.PoolState.Buffer.WriteByte('/')
        fs.PoolState.Buffer.WriteInt(int64(month))
        fs.PoolState.Buffer.WriteByte('/')
        fs.PoolState.Buffer.WriteInt(int64(day))
    }

    if fs.WithTime {
        fs.PoolState.Buffer.WriteSpace()
        fs.PoolState.Buffer.WriteInt(int64(now.Hour()))
        fs.PoolState.Buffer.WriteByte(':')
        fs.PoolState.Buffer.WriteInt(int64(now.Minute()))
        fs.PoolState.Buffer.WriteByte(':')
        fs.PoolState.Buffer.WriteInt(int64(now.Second()))
    }

    if fs.WithNS {
        fs.PoolState.Buffer.WriteByte('.')
        fs.PoolState.Buffer.WriteInt(int64(now.Nanosecond()))
    }
}

func writePID(fs *formatters.FormatterConfig)  {
    if fs.WithPID {
        fs.PoolState.Buffer.WriteSpace()
        fs.PoolState.Buffer.WriteByte('[')
        fs.PoolState.Buffer.WriteInt(int64(os.Getpid()))
        fs.PoolState.Buffer.WriteByte(']')
    }
}
