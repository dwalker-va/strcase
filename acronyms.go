package strcase

// ConfigureAcronym allows you to add additional words which will be considered acronyms
func ConfigureAcronym(key, val string) {
	defaultConverter.ConfigureAcronym(key, val)
}

func (c *Converter) ConfigureAcronym(key, val string) {
	c.lock.Lock()
	c.uppercaseAcronym[key] = val
	c.lock.Unlock()
}
