package pool

// Get is a type-safe wrapper of the sync.Pool.Get()
func (p *Pool) Get() *State {
    return p.p.Get().(*State)
}

// Set is a type-safe wrapper of the sync.Pool.Put()
func (p *Pool) Set(ps *State) {
    p.p.Put(ps)
}

// Free resets the pool state to a starting point and
// puts it into the pool.
func (p *Pool) Free(ops *State) {
    ops.Buffer.Reset()
    ops.Map.Reset()
    p.Set(ops)
}
