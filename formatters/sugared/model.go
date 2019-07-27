package sugared

import "github.com/olehan/kek/formatters"

type (
    SugaredFormatter struct {
        *SugaredBaseFormatter
        *SugaredTemplateFormatter
        *SugaredStructuredFormatter
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
