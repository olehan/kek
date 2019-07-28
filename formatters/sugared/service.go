package sugared

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
    bold  = colors.Bold.String()
)

func WriteMetaInfo(fs *formatters.FormatterConfig) {
    WriteName(fs, " | ")

    WriteLevel(fs)

    if fs.WithColors {
        fs.PoolState.Buffer.WriteString(bold)
    }

    fs.PoolState.Buffer.WriteSpace()

    WriteDateTime(fs)

    if fs.WithColors {
        fs.PoolState.Buffer.WriteString(reset)
    }

    WritePID(fs)
    fs.PoolState.Buffer.WriteString(":   ")
}

func WriteName(fs *formatters.FormatterConfig, separator string) {
    if len(fs.Name) > 0 {
        shouldWriteColor := fs.WithColors && len(fs.NameColor) > 0

        if shouldWriteColor {
            fs.PoolState.Buffer.WriteString(fs.NameColor)
        }

        fs.PoolState.Buffer.WriteString(fs.Name)

        WriteNameTabulation(fs)

        if len(separator) > 0 {
            fs.PoolState.Buffer.WriteString(separator)
        }

        if shouldWriteColor {
            fs.PoolState.Buffer.WriteString(reset)
        }
    }
}

func WriteNameTabulation(fs *formatters.FormatterConfig) {
    if fs.WithNameTabulation {
        for i := 0; i < names.LongestNameLen() - len(fs.Name); i++ {
            println("has", string(fs.PoolState.Buffer))
            fs.PoolState.Buffer.WriteSpace()
        }
    }
}

func WriteLevel(fs *formatters.FormatterConfig) {
    if fs.WithColors {
        fs.PoolState.Buffer.WriteString(levels.ColoredLevelMap[fs.Level])
    } else {
        fs.PoolState.Buffer.WriteString(levels.NonColoredLevelMap[fs.Level])
    }

    WriteLevelTabulation(fs)
}

func WriteLevelTabulation(fs *formatters.FormatterConfig) {
    for i := 0; i < levels.LongestLevelLen() - len(levels.NonColoredLevelMap[fs.Level]); i++ {
        fs.PoolState.Buffer.WriteSpace()
    }
}

func WriteDateTime(fs *formatters.FormatterConfig) {
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

func WritePID(fs *formatters.FormatterConfig) {
    if fs.WithPID {
        fs.PoolState.Buffer.WriteSpace()
        fs.PoolState.Buffer.WriteByte('[')
        fs.PoolState.Buffer.WriteInt(int64(os.Getpid()))
        fs.PoolState.Buffer.WriteByte(']')
    }
}
