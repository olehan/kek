package pool

import "testing"

var (
    _testPool = NewPool()
)

func TestPool(t *testing.T) {
    s := _testPool.Get()
    s.Buffer.WriteSpace()
    _testPool.Set(s)
    s = _testPool.Get()
    if string(s.Buffer) != " " {
        panic("Buffer from state has no assigned space")
    }
    _testPool.Free(s)
}
