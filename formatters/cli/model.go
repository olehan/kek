package cli

import (
    "github.com/olehan/kek/formatters"
    "strings"
)

type (
    FormatterModel struct {
        formatters.Formatter
        *formatters.FormatterUtils
        appName string
    }
)

var (
    Formatter = NewCLIFormatter()
)

func NewCLIFormatter(appName ...string) formatters.Formatter {
    return &FormatterModel{
        FormatterUtils: formatters.NewFormatterUtils(),
        appName: strings.Join(appName, "."),
    }
}
