package suggared

import "github.com/olehan/kek/formatters"

type (
    SugaredBaseFormatter struct {
        formatters.BaseFormatter
        *formatters.FormatterUtils
    }
)

func NewSugaredBaseFormatter(fu ...*formatters.FormatterUtils) *SugaredBaseFormatter {
    return &SugaredBaseFormatter{
        FormatterUtils:  formatters.TakeFirstUtil(fu...),
    }
}
