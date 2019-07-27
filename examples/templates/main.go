package main

import (
    "github.com/olehan/kek"
    "github.com/olehan/kek/ds"
)

func main() {
    logger := kek.NewLogger()

    logger.Debug.
        PrintT("wassup, {} - {}", "boi", 123).
        PrintTM("wassup, NewMap{{name}}, Age: {{age}}",
            ds.NewMap().
                Set("name", "Boi").
                Set("age", 17),
        ).
        PrintTM("wassup, Map{{name}}", ds.Map{
            "name": "Boi",
        }).
        PrintTKV("wassup, KeyValue{{name}} {{34}} {{134}} - {{unknown}}",
            "name", "Boi",
            "34", true,
            // Skips the key that is not a string.
            134, "ooh",
            // Skips the key value pair without a value.
            "unknown",
        )
}
