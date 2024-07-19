package tuid

import "time"

type GenOption interface {
	apply(*TUID)
}

type funcGenOption struct {
	f func(*TUID)
}

func (fgo *funcGenOption) apply(id *TUID) {
	fgo.f(id)
}

func newFuncDialOption(f func(*TUID)) *funcGenOption {
	return &funcGenOption{
		f: f,
	}
}

func WithPrefix(prefix string) GenOption {
	return newFuncDialOption(func(id *TUID) {
		id.prefix = prefix
	})
}

func WithTimeLayout(layout string) GenOption {
	return newFuncDialOption(func(id *TUID) {
		id.timeLayout = layout
	})
}

func WithTime(t time.Time) GenOption {
	return newFuncDialOption(func(id *TUID) {
		id.time = t
	})
}

func WithRandomCode(length int) GenOption {
	return newFuncDialOption(func(id *TUID) {
		id.randomCode = RandomString(length)
	})
}

func WithSequence(seq int64) GenOption {
	return newFuncDialOption(func(id *TUID) {
		id.seq = seq
	})
}
