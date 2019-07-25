package suggared

import "github.com/olehan/kek/formatters"

type (
    SugaredStructuredFormatter struct {
        formatters.StructuredFormatter
        *formatters.FormatterUtils
    }
)

func NewSugaredStructuredFormatter(fu ...*formatters.FormatterUtils) *SugaredStructuredFormatter {
    return &SugaredStructuredFormatter{
        FormatterUtils: formatters.TakeFirstUtil(fu...),
    }
}
