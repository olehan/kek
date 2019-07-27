package main

import (
    "github.com/olehan/kek"
    "github.com/olehan/kek/formatters/minified"
)

func main() {
    logger := kek.NewLogger().SetFormatter(minified.Formatter)
    logger.Debug.Print("hello, kek!")
}
