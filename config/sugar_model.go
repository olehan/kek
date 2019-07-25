package config

type (
    SugarConfig struct {
        WithColors bool
        WithTime   bool
        WithPID    bool
    }
)

func NewSugarConfig() *SugarConfig {
    return &SugarConfig{}
}
