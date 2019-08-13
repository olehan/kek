package main

import (
    "github.com/olehan/kek"
    "github.com/olehan/kek/formatters/cli"
)

func main() {
    logger := kek.NewLogger("main").SetFormatter(cli.NewCLIFormatter("rise"))

    logger.Info.Println("Hello, World!")
    logger.Debug.Println("Hello, World!")
    logger.Warn.Println("Hello, World!")
    logger.Note.Println("Hello, World!")
    logger.Error.Println("Hello, World!")
    logger.Succ.Println("Hello, World!")
}
