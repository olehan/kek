package sugar

func (m Map) Reset() {
    for k := range m {
        m[k] = nil
    }
}
