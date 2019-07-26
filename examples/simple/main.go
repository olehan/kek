package main

import "github.com/olehan/kek"

func main() {
    logger := kek.NewLogger()
    logger.Info.Println("Hello, World!")
}
