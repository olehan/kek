package sugared

import "github.com/olehan/kek/formatters"

type (
    SugaredFormatter struct {
        formatters.Formatter
        *formatters.FormatterUtils
    }
)

var (
    Formatter = NewSugaredFormatter()
)

func NewSugaredFormatter() formatters.Formatter {
    return &SugaredFormatter{
        FormatterUtils: formatters.NewFormatterUtils(),
    }
}
