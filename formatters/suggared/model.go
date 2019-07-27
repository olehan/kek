package suggared

import "github.com/olehan/kek/formatters"

type (
    SugaredFormatter struct {
        *SugaredBaseFormatter
        *SugaredTemplateFormatter
        *SugaredStructuredFormatter

        formatterUtils *formatters.FormatterUtils
    }
)

var (
    Formatter = NewSugaredFormatter()
)

func NewSugaredFormatter() formatters.Formatter {
    formatterUtils := formatters.NewFormatterUtils()
    return &SugaredFormatter{
        SugaredBaseFormatter: NewSugaredBaseFormatter(formatterUtils),
        SugaredTemplateFormatter: NewSugaredTemplateFormatter(formatterUtils),
        SugaredStructuredFormatter: NewSugaredStructuredFormatter(formatterUtils),
    }
}
