package importer

import (
	"github.com/tableauio/tableau/internal/importer/book"
	"github.com/tableauio/tableau/options"
)

type Options struct {
	Sheets []string              // sheet names to import
	Parser book.SheetParser      // parser to parse the worksheet
	Header *options.HeaderOption // header settings.
}

// Option is the functional option type.
type Option func(*Options)

func Sheets(sheets []string) Option {
	return func(opts *Options) {
		opts.Sheets = sheets
	}
}

func Parser(parser book.SheetParser) Option {
	return func(opts *Options) {
		opts.Parser = parser
	}
}

func Header(header *options.HeaderOption) Option {
	return func(opts *Options) {
		opts.Header = header
	}
}

func newDefaultOptions() *Options {
	return &Options{}
}

func parseOptions(setters ...Option) *Options {
	// Default Options
	opts := newDefaultOptions()
	for _, setter := range setters {
		setter(opts)
	}
	return opts
}
