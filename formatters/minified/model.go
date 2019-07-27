package minified

import "github.com/olehan/kek/formatters"

type (
    MinifiedFormatter struct {
        formatters.Formatter
        *formatters.FormatterUtils
    }
)

var (
    Formatter = NewMinifiedFormatter()
)

func NewMinifiedFormatter() formatters.Formatter {
    return &MinifiedFormatter{
        FormatterUtils: formatters.NewFormatterUtils(),
    }
}
