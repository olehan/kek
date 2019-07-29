package pool

import (
    "github.com/olehan/kek/buffer"
    "github.com/olehan/kek/ds"
    "sync"
)

type (
    // Pool is a wrapper above the sync.Pool that gives
    // type-safe Get and Set functions.
    Pool struct {
        p *sync.Pool
    }

    // State is a value saved into a New func of the Pool.
    State struct {
        Buffer buffer.Buffer
        Map    ds.Map
    }
)

// NewPool returns a new Pool wrapper.
func NewPool() *Pool {
    return &Pool{
        p: &sync.Pool{
            New: func() interface{} {
                return &State{
                    Buffer: buffer.NewBuffer(),
                    Map:    ds.Map{},
                }
            },
        },
    }
}
