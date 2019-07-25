package config

import "github.com/olehan/kek/pool"

type (
    PoolConfig struct {
        Pool *pool.Pool
        PoolState *pool.PoolState
    }
)

func NewPoolConfig() *PoolConfig {
    return &PoolConfig{}
}

func (p *PoolConfig) SetPool(pool *pool.Pool) *PoolConfig {
    p.Pool = pool
    return p
}

func (p *PoolConfig) SetPoolState(ps *pool.PoolState) *PoolConfig {
    p.PoolState = ps
    return p
}

func (p *PoolConfig) Value() PoolConfig {
    return *p
}

func (p *PoolConfig) Copy() *PoolConfig {
    cp := *p
    return &cp
}

func (p *PoolConfig) CopyValue() PoolConfig {
    return p.Copy().Value()
}
