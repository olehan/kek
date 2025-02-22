package levels

// LongestLevelLen returns a copy of the longestLevelLen.
func LongestLevelLen() int {
    return longestLevelLen
}

func getLongestLevelLen(levelMap LevelMap) int {
    temp := 0
    for _, n := range levelMap {
        nLen := len(n)
        if temp < nLen {
            temp = nLen
        }
    }
    return temp
}
