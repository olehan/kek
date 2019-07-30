package pool

import "testing"

var (
    _testPool = NewPool()
)

func TestPool(t *testing.T) {
    s := _testPool.Get()
    s.Buffer.WriteSpace()
    _testPool.Set(s)
    if string(s.Buffer) != " " {
        panic("Buffer from state has no assigned space")
    }
    _testPool.Free(s)
    if len(s.Buffer) != 0 {
        panic("Buffer hasn't changed after the reset")
    }
}
