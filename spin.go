// Package spin implements a simple console spinner. See the tests and examples for examples.
package spin

import (
	"fmt"
	"io"
	"os"
	"time"
)

var standard = New()

// GoStart forwards to a package-level *Spinner (for convenience).
func GoStart() {
	go Start()
}

// Start forwards to a package-level *Spinner (for convenience).
func Start() {
	standard.Start()
}

// Stop forwards to a package-level *Spinner (for convenience).
func Stop() {
	standard.Stop()
}

type Style string

var (
	StyleLine     = Style("|/-\\")
	StylePops     = Style("-=*%*=")
	StyleSteps    = Style("▁▃▄▅▆▇█▇▆▅▄▃")
	StyleShutter  = Style("▉▊▋▌▍▎▏▎▍▌▋▊▉")
	StyleBrackets = Style(">})|({<-<{(|)}>")
	StyleNumbers  = Style("0123456789")
	StyleAlphabet = Style("abcdefghijklmnopqrstuvwxyz")
)

// Spinner prints a repeating pattern to os.Stdout by printing a sequence of characters
// interspersed with carriage returns. A Spinner is controlled by the provided methods:
// Start, GoStart (like calling `go Start()`, and Stop.
type Spinner struct {
	out    *output
	style  Style
	delay  time.Duration
	prefix string
	suffix string
	stop   chan struct{}
}

// New creates a spinner which you can start and stop.
func New(options ...option) *Spinner {
	s := &Spinner{stop: make(chan struct{})}
	for _, option := range append(defaults, options...) {
		option(s)
	}
	return s
}

// Start begins the spinner on the current goroutine (hopefully you've got another goroutine that can call Stop...).
func (self *Spinner) Start() {
	for {
		select {
		case <-self.stop:
			_, _ = fmt.Fprint(self.out, "\r") // erase any residual spinner markings
			return
		default:
			self.spinCycle()
		}
	}
}
func (self *Spinner) spinCycle() {
	for _, symbol := range self.style {
		_, _ = fmt.Fprintf(self.out, "\r%s%s%s", self.prefix, string(symbol), self.suffix)
		time.Sleep(self.delay)
	}
}

// Stop sends a signal to stop the spinner.
func (self *Spinner) Stop() {
	self.stop <- struct{}{}
}

////////////////////////////////////////////////////////

type output struct {
	out io.Writer
}

func (self *output) Write(p []byte) (int, error) {
	if self == nil {
		return os.Stdout.Write(p)
	}
	return self.out.Write(p)
}
