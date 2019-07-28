package main

import (
    "github.com/olehan/kek"
    "github.com/olehan/kek/formatters/minified"
)

func main() {
    logger := kek.NewLogger().SetFormatter(minified.Formatter)
    logger.Debug.
        PrintSKV("message for key values",
            "key1", "value1",
            "key2", 242456246,
            "key3", true,
            "key4", 135135.13413,
        )
}
