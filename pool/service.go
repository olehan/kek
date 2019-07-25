package pool

func (p *Pool) Get() *PoolState {
    return p.p.Get().(*PoolState)
}

func (p *Pool) Set(ps *PoolState) {
    p.p.Put(ps)
}

func (p *Pool) Free(ops *PoolState) {
    ops.Buffer.Reset()
    ops.Map.Reset()
    p.Set(ops)
}
