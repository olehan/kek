package sugar

func (m Map) Set(key string, value interface{}) Map {
    m[key] = value
    return m
}

func (m Map) Delete(key string) Map {
    m[key] = nil
    return m
}

func (m Map) Reset() {
    for k := range m {
        m[k] = nil
    }
}
