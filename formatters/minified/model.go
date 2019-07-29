package minified

import "github.com/olehan/kek/formatters"

type (
    // FormatterModel is struct that implements formatter for
    // the sugared format.
    FormatterModel struct {
        formatters.Formatter
        *formatters.FormatterUtils
    }
)

var (
    // Formatter is a singleton for a minified formatter instance.
    Formatter = NewMinifiedFormatter()
)

// NewSugaredFormatter returns a new minified formatter.
func NewMinifiedFormatter() formatters.Formatter {
    return &FormatterModel{
        FormatterUtils: formatters.NewFormatterUtils(),
    }
}
