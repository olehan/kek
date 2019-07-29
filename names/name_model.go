package names

var (
    longestNameLen = 0
)

// LongestNameLen returns a copy of the longestNameLen.
func LongestNameLen() int {
    return longestNameLen
}

// SetLongestNameLen sets the new longest name length
// if its longer then the current one.
func SetLongestNameLen(l int) int {
    if longestNameLen < l {
        longestNameLen = l
        return l
    }
    return longestNameLen
}
