package pool

import (
    "github.com/olehan/kek/buffer"
    "github.com/olehan/kek/ds"
    "sync"
)

type (
    Pool struct {
        p *sync.Pool
    }

    PoolState struct {
        Buffer buffer.Buffer
        Map    ds.Map
    }
)

func NewPool() *Pool {
    return &Pool{
        p: &sync.Pool{
            New: func() interface{} {
                return &PoolState{
                    Buffer: buffer.NewBuffer(),
                    Map:    ds.Map{},
                }
            },
        },
    }
}
