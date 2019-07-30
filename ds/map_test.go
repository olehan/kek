package ds

import "testing"

func TestMap(t *testing.T) {
    m := NewMap()
    if len(m) != 0 {
        panic("map is not empty")
    }

    expectedKey := "key1"
    expectedValue := "value1"

    m.Set(expectedKey, expectedValue)

    if m[expectedKey] != expectedValue {
        panic("map has no key " + expectedKey)
    }

    m.Delete(expectedKey)

    if len(m) != 0 {
        panic("map is not empty")
    }

    m.Set(expectedKey, expectedValue)
    m.Set(expectedValue, expectedKey)

    m.Reset()

    if len(m) != 0 {
        panic("map is not empty")
    }
}
