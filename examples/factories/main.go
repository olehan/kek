package main

import (
    "github.com/olehan/kek"
    "github.com/olehan/kek/formatters/sugared"
    "os"
)

func main() {
    // Creating a new factory to populate new logger with a specific configuration.
    factory := kek.NewFactory(os.Stdout, sugared.Formatter, "factory", "name")
    // Setting configurations of the factory.
    factory.
        SetWithDateTime(false, true, false).
        SetWithPID(false)

    // Now logger extends the factory configurations...
    logger := factory.NewLogger("logger", "name")
    // ...but if you want a random color for every logger
    // you still need to execute SetRandomNameColor for
    // every created logger individually.
    logger.SetRandomNameColor()
    logger.Debug.Println("wassup, kek")
    logger.Info.Println("info works", 2)
}
