package ds

type (
    // Map is representation of a string keyed and interface valued
    // map used in several printers as an easiest data structured
    // to work with.
    Map map[string]interface{}
)

// NewMap return a new Map.
func NewMap() Map {
    return Map{}
}
