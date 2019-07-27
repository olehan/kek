package names

var (
    longestNameLen = 0
)

func LongestNameLen() int {
    return longestNameLen
}

func SetLongestNameLen(l int) int {
    if longestNameLen < l {
        longestNameLen = l
        return l
    }
    return longestNameLen
}
