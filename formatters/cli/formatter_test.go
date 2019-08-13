package cli

import (
    "github.com/olehan/kek/config"
    "github.com/olehan/kek/ds"
    "github.com/olehan/kek/formatters"
    "github.com/olehan/kek/levels"
    "github.com/olehan/kek/names"
    "github.com/olehan/kek/pool"
    "os"
    "testing"
)

var (
    _testFormatter       = NewCLIFormatter()
    _testFormatterConfig = formatters.NewFormatterConfig(
        config.NewConfig().
            SetWithColors(true).
            SetName("name").
            SetRandomNameColor().
            SetWriter(os.Stdout)).
        SetLevel(levels.Info).
        SetPoolState(pool.NewPool().Get())
)

func init() {
    names.SetLongestNameLen(10)
}

func TestFormatterModel_Print(t *testing.T) {
    _testFormatter.Print(_testFormatterConfig, "value")
}

func TestFormatterModel_PrintTemplate(t *testing.T) {
    _testFormatter.PrintTemplate(_testFormatterConfig, "template {}", "value")
}

func TestFormatterModel_PrintTemplateMap(t *testing.T) {
    _testFormatter.PrintTemplateMap(
        _testFormatterConfig,
        "template {{key1}}",
        ds.NewMap().Set("key1", "value"),
    )
}

func TestFormatterModel_PrintTemplateKeyValue(t *testing.T) {
    _testFormatter.PrintTemplateKeyValue(
        _testFormatterConfig,
        "template {{key1}}",
        "key1", "value",
        135, "skip",
        "skip",
    )
}

func TestFormatterModel_PrintStructKeyValues(t *testing.T) {
    _testFormatter.PrintStructKeyValues(
        _testFormatterConfig,
        "message",
        "key1", "value1",
        "key2", "value2",
        "key3", "value3",
        13413, "1341",
        "skip",
    )
}
