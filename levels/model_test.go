package levels

import "testing"

func TestLevel(t *testing.T) {
    t.Log(ColoredLevelMap[Info], "info level")
    t.Log(ColoredLevelMap[Succ], "success level")
    t.Log(ColoredLevelMap[Debug], "debug level")
    t.Log(ColoredLevelMap[Note], "note level")
    t.Log(ColoredLevelMap[Warn], "warn level")
    t.Log(ColoredLevelMap[Error], "error level")
    t.Log(ColoredLevelMap[Fatal], "fatal level")
    t.Log(ColoredLevelMap[Panic], "panic level")
}
