package zcache

import "time"

type Options struct {
	Addr     string
	Port     int
	CleanSeq time.Duration
	LandSeq  time.Duration
}

type OptionFunc func(opts *Options)

func loadOptions(options ...OptionFunc) *Options {
	opts := new(Options)
	for _, option := range options {
		option(opts)
	}
	return opts
}

func WithOptions(options Options) OptionFunc {
	return func(opts *Options) {
		*opts = options
	}
}

func WithAddr(Addr string) OptionFunc {
	return func(opts *Options) {
		opts.Addr = Addr
	}
}

func WithPort(Port int) OptionFunc {
	return func(opts *Options) {
		opts.Port = Port
	}
}

func WithCleanSeq(CleanSeq time.Duration) OptionFunc {
	return func(opts *Options) {
		opts.CleanSeq = CleanSeq
	}
}

func WithLandSeq(LandSeq time.Duration) OptionFunc {
	return func(opts *Options) {
		opts.LandSeq = LandSeq
	}
}
