package spin

import "time"

var defaults = []option{
	Options.Style(StyleLine),
	Options.Delay(time.Millisecond * 100),
	Options.Prefix(""),
	Options.Suffix(""),
}

type option func(*Spinner)

type options struct{}

var Options options

func (options) Style(v Style) option         { return func(s *Spinner) { s.style = v } }
func (options) Delay(v time.Duration) option { return func(s *Spinner) { s.delay = v } }
func (options) Prefix(v string) option       { return func(s *Spinner) { s.prefix = v } }
func (options) Suffix(v string) option       { return func(s *Spinner) { s.suffix = v } }
