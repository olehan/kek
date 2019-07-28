package formatters

import "github.com/olehan/kek/ds"

type (
    Formatter interface {
        BaseFormatter
        TemplateFormatter
        StructuredFormatter
    }

    BaseFormatter interface {
        // value 1 v true 1.1
        Print(fs *FormatterConfig, values ...interface{})
    }

    TemplateFormatter interface {
        PrintTemplate(fs *FormatterConfig, template string, values ...interface{})
        // PrintFormatMap should print values that given in the form of a map.
        // Expected usage:
        //  format: "values {{ a }} - {{ b }} - {{ unknown }}"
        //  values: Map{ "a": "it is A", "b": "rush B", }
        //  output: "values it is A - rush B - {{ unknown }}"
        PrintTemplateMap(fs *FormatterConfig, template string, values ds.Map)
        // PrintKeyValue should print values that given in the form of interface arguments
        // that is represented as key and value pair.
        // Expected usage:
        //  format: "values {{ a }} - {{ b }} - {{ unknown }}"
        //  values: []interface{}{ "a", "it is A", "b", "rush B" }
        //  output: "values it is A - rush B - {{ unknown }}"
        PrintTemplateKeyValue(fs *FormatterConfig, template string, keyValues ...interface{})
    }

    StructuredFormatter interface {
        PrintStructKeyValues(fs *FormatterConfig, message string, keyValues ...interface{})
    }
)
