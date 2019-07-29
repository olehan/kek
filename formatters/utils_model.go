package formatters

type (
    // FormatterUtils is a struct that has a useful set of
    // formatting util functions that may help working with
    // interface{} value.
    FormatterUtils struct {
    }
)

// NewFormatterUtils returns new formatter utils.
func NewFormatterUtils() *FormatterUtils {
    return &FormatterUtils{}
}

// TakeFirstUtil takes the first formatter util if any presented
// and returns it, otherwise will create a new one.
func TakeFirstUtil(fu ...*FormatterUtils) *FormatterUtils {
    var formatterUtils *FormatterUtils
    if len(fu) > 0 {
        formatterUtils = fu[0]
    } else {
        formatterUtils = NewFormatterUtils()
    }
    return formatterUtils
}
