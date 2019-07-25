package formatters

import "github.com/olehan/kek/config"

type (
    FormatterConfig struct {
        *config.PoolConfig
        *config.BaseConfig
        *config.SugarConfig
    }
)

func NewFormatterConfig(
    pc *config.PoolConfig,
    bc *config.BaseConfig,
    sc *config.SugarConfig,
) *FormatterConfig {
    return &FormatterConfig{
        PoolConfig: pc,
        BaseConfig: bc,
        SugarConfig: sc,
    }
}
