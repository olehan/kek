package benchmarks

import (
    "github.com/olehan/kek"
    "github.com/olehan/kek/ds"
    "io"
    "os"
)

type emptyWriter struct {
    io.Writer
}

func (emptyWriter) Write(p []byte) (n int, err error) {
    return 0, err
}

var (
    _testValues       = []interface{}{257314.234, '-', false}
    _testMapValues    = ds.NewMap().
        Set("size", _testValues[0]).
        Set("separator", _testValues[1]).
        Set("visibility", _testValues[2])

    _testSimpleMock   = []interface{}{
        "UserModel update success, size:",
        _testValues[0],
        _testValues[1],
        "Is Visible:",
        _testValues[2],
    }

    _testSimpleFormat = "UserModel update success, size: {} {} Is Visible: {}"
    _testMapFormat    = "UserModel update success, size: {{size}} {{separator}} Is Visible: {{visibility}}"
)

var (
    _testEmptyLogger = kek.NewLogger("very long test name with empty writer")
    _testStdoutLogger = kek.NewLogger("very long test name with stdout writer")
)

func init() {
    _testEmptyLogger.SetWriter(emptyWriter{})
    _testStdoutLogger.SetWriter(os.Stdout)
}
