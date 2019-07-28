package main

import (
    "github.com/olehan/kek"
    "github.com/olehan/kek/formatters/minified"
    "github.com/olehan/kek/formatters/sugared"
    "os"
)

func main() {
    factory := kek.NewFactory(os.Stdout, minified.Formatter)

    minifiedLogger := factory.NewLogger("name")
    minifiedLogger.SetRandomNameColor()
    minifiedLogger.Debug.Println("hello, kek!")
    minifiedLogger.Debug.PrintSKV("here goes structures",
        "url", "http://127.0.0.1:1110",
        "jwt", "v75onlb3iv3i7t5",
        "salt", 123,
    )

    sugaredLogger := factory.NewLogger("name").SetFormatter(sugared.Formatter)
    sugaredLogger.SetRandomNameColor()
    sugaredLogger.Info.Println("hello, not minified logger")
    sugaredLogger.Info.PrintSKV("cache state",
        "keys", 92334,
        "misses", 573,
        "matches", 24247754,
        "ttl", 1300.1220,
        "active", true,
    )
}
