package ds

// Set sets a new value into the Map.
func (m Map) Set(key string, value interface{}) Map {
    m[key] = value
    return m
}

// Delete removes value from the Map by setting nil value
// by the given key.
func (m Map) Delete(key string) Map {
    delete(m, key)
    return m
}

// Reset resets the whole map by setting nil value to
// every key init.
func (m Map) Reset() Map {
    for k := range m {
        delete(m, k)
    }
    return m
}
