package sugared

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
    // Formatter is a singleton for a sugared formatter instance.
    Formatter = NewSugaredFormatter()
)

// NewSugaredFormatter returns a new sugared formatter.
func NewSugaredFormatter() formatters.Formatter {
    return &FormatterModel{
        FormatterUtils: formatters.NewFormatterUtils(),
    }
}
