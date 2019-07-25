package colors

// Color represent ansi escape sequence code.
type Color uint8

const (
    // Black is ansi code that represents Black color.
    Black Color = iota + 30
    // Red is ansi code that represents Red color.
    Red
    // Green is ansi code that represents Green color.
    Green
    // Yellow is ansi code that represents Yellow color.
    Yellow
    // Blue is ansi code that represents Blue color.
    Blue
    // Magenta is ansi code that represents Magenta color.
    Magenta
    // Cyan is ansi code that represents Cyan color.
    Cyan
    // White is ansi code that represents White color.
    White
)

const (
    // Reset is ansi code that resets the sequence.
    Reset Color = iota
    // Bold is ansi code that gives a text bold weight.
    Bold
    // Faint is ansi code that gives a text light weight.
    Faint
    // Italic is ansi code that gives a text italic style.
    Italic
    // Underline is ansi code that puts an underline for the text.
    Underline
    // SlowBlink is ansi code that gives slow blink effect
    // for the text.
    SlowBlink
    // RapidBlink is ansi code that gives rapid blink effect
    // for the text.
    RapidBlink
    // Reverse is ansi code that reverses the text.
    Reverse
)
