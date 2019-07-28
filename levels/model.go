package levels

import "github.com/olehan/kek/colors"

type (
    // Level is an int that represents the level code.
    Level int8
    // LevelMap is a map that contains a string
    // representation of level by its code.
    LevelMap map[Level]string
)

const (
    // Info is code for INFO level.
    Info Level = iota
    // Succ is code for SUCC level.
    // It's Success, btw ヽ(￣ω￣ )
    Succ
    // Debug is code for DEBUG level.
    Debug
    // Note is code for NOTE level.
    Note
    // Warn is code for WARN level.
    Warn
    // Error is code for ERROR level.
    Error
    // Fatal is code for FATAL level.
    Fatal
    // Panic is code for PANIC level.
    Panic
)

var (
    longestLevelLen = getLongestLevelLen(NonColoredLevelMap)
    // NonColoredLevelMap is a level map of non colored level names.
    NonColoredLevelMap = LevelMap{
        Info:  "info",
        Succ:  "succ",
        Debug: "debug",
        Note:  "note",
        Warn:  "warn",
        Error: "error",
        Fatal: "fatal",
        Panic: "panic",
    }
    // ColoredLevelMap is a level map of colored level names.
    ColoredLevelMap = LevelMap{
        Info: colors.Blue.Add(NonColoredLevelMap[Info]),
        Succ: colors.Green.Add(NonColoredLevelMap[Succ], colors.Bold),
        Debug: colors.Magenta.Add(NonColoredLevelMap[Debug]),
        Note: colors.Cyan.Add(NonColoredLevelMap[Note], colors.Bold),
        Warn: colors.Yellow.Add(NonColoredLevelMap[Warn], colors.Bold),
        Error: colors.Red.Add(NonColoredLevelMap[Error]),
        Fatal: colors.Red.Add(NonColoredLevelMap[Note], colors.Bold),
        Panic: colors.Red.Add(NonColoredLevelMap[Panic], colors.Bold, colors.Underline),
    }
)
