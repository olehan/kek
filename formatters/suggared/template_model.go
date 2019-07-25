package suggared

import "github.com/olehan/kek/formatters"

type (
    SugaredTemplateFormatter struct {
        formatters.TemplateFormatter
        *formatters.FormatterUtils
    }
)

func NewSugaredTemplateFormatter(fu ...*formatters.FormatterUtils) *SugaredTemplateFormatter {
    return &SugaredTemplateFormatter{
        FormatterUtils:  formatters.TakeFirstUtil(fu...),
    }
}
