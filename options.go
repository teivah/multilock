package multilock

type Option interface {
	apply(*options)
}

type options struct {
	distribution func(s string, length int) int
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

func WithCustomDistribution(hash func(s string, length int) int) Option {
	return newFuncOption(func(options *options) {
		options.distribution = hash
	})
}
