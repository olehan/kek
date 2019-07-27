package main

import "github.com/olehan/kek"

func main() {
    logger := kek.NewLogger()

    logger.Info.Println("Hello, World!")
    logger.Debug.Println("Hello, World!")
    logger.Warn.Println("Hello, World!")
    logger.Note.Println("Hello, World!")
    logger.Error.Println("Hello, World!")
    logger.Panic.Println("Hello, World!")
    logger.Succ.Println("Hello, World!")
}
