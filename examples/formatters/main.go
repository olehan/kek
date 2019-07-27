package main

import (
    "github.com/olehan/kek"
    "github.com/olehan/kek/formatters/minified"
    "github.com/olehan/kek/formatters/sugared"
)

func main() {
    minifiedLogger := kek.NewLogger("name").SetFormatter(minified.Formatter)
    minifiedLogger.SetRandomNameColor()
    minifiedLogger.Debug.Println("hello, kek!")

    sugaredLogger := kek.NewLogger("name").SetFormatter(sugared.Formatter)
    sugaredLogger.SetRandomNameColor()
    sugaredLogger.Info.Println("hello, not minified logger")
}
