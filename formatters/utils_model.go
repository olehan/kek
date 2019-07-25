package formatters

type (
    FormatterUtils struct {
    }
)

func NewFormatterUtils() *FormatterUtils {
    return &FormatterUtils{}
}

func TakeFirstUtil(fu ...*FormatterUtils) *FormatterUtils {
    var formatterUtils *FormatterUtils
    if len(fu) > 0 {
        formatterUtils = fu[0]
    } else {
        formatterUtils = NewFormatterUtils()
    }
    return formatterUtils
}
