package levels

import "testing"

func TestLongestLevelLen(t *testing.T) {
    l := LongestLevelLen()
    if l != len(NonColoredLevelMap[Debug]) {
        panic("LongestLevelLen has unexpected value")
    }
}
