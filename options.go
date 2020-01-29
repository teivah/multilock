package multilock

// Option contains the multilock options
type Option interface {
	apply(*options)
}

type options struct {
	distribution func(i interface{}, length int) int
}

type funcOption struct {
	f func(*options)
}

func (fdo *funcOption) apply(do *options) {
	fdo.f(do)
}

func newFuncOption(f func(*options)) *funcOption {
	return &funcOption{
		f: f,
	}
}

// WithCustomDistribution allows to specify a custom distribution function
func WithCustomDistribution(distribution func(i interface{}, length int) int) Option {
	return newFuncOption(func(options *options) {
		options.distribution = distribution
	})
}
