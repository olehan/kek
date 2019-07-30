package names

import "testing"

func TestLongestNameLen(t *testing.T) {
    if LongestNameLen() != 0 {
        panic("there isn't init value of the LongestNameLen")
    }

    expectedValue := 4

    SetLongestNameLen(expectedValue)

    if LongestNameLen() != expectedValue {
        panic("LongestNameLen is not the expected value")
    }

    SetLongestNameLen(expectedValue - 1)

    if LongestNameLen() != expectedValue {
        panic("LongestNameLen is not the expected value")
    }
}
