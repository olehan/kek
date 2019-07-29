package formatters

import "github.com/olehan/kek/ds"

type (
    // Formatter is an interface that summarize formatter interfaces.
    Formatter interface {
        BaseFormatter
        TemplateFormatter
        StructuredFormatter
    }

    // BaseFormatter is an interface that describes
    // base formatting functionality.
    BaseFormatter interface {
        Print(fs *FormatterConfig, values ...interface{})
    }

    // TemplateFormatter is an interface that describes
    // template formatting functionality.
    TemplateFormatter interface {
        // PrintTemplate should print values by the given format/template.
        // Expected usage:
        //  format: "values {} - {} - {}"
        //  values: []interface{}{ "a", "it is A", "b", "rush B" }
        //  output: "values it is A - rush B - {}"
        PrintTemplate(fs *FormatterConfig, template string, values ...interface{})
        // PrintFormatMap should print values that given in the form of a map
        // according to the format/template.
        //
        // Expected usage:
        //  format: "values {{a}} - {{b}} - {{unknown}}"
        //  values: Map{ "a": "it is A", "b": "rush B", }
        //  output: "values it is A - rush B - {{ unknown }}"
        PrintTemplateMap(fs *FormatterConfig, template string, values ds.Map)
        // PrintKeyValue should print values that given in the form of interface arguments
        // that is represented as key and value pair according to the format/template.
        // Expected usage:
        //  format: "values {{a}} - {{b}} - {{unknown}}"
        //  values: []interface{}{ "a", "it is A", "b", "rush B" }
        //  output: "values it is A - rush B - {{ unknown }}"
        PrintTemplateKeyValue(fs *FormatterConfig, template string, keyValues ...interface{})
    }

    // StructuredFormatter is an interface that describes
    // structured formatting functionality.
    StructuredFormatter interface {
        PrintStructKeyValues(fs *FormatterConfig, message string, keyValues ...interface{})
    }
)
