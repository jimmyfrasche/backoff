package backoff

//A Factory is a helper field for configuration that need to generate new Strategies.
type Factory func() Strategy

//New invokes f.
//If f or the returned Strategy is nil,
//then New returns None().
func (f Factory) New() Strategy {
	if f == nil {
		return None()
	}
	s := f()
	if s == nil {
		s = None()
	}
	return s
}

//NewOrDefault invokes f.
//If f or the returned Strategy is nil,
//then NewOrDefault returns Default.New().
func (f Factory) NewOrDefault(Default Factory) Strategy {
	if f == nil {
		return Default.New()
	}
	s := f.New()
	if s == nil {
		return Default.New()
	}
	return s
}
