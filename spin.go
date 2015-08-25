// Package spin implements a simple console spinner. See the tests and examples for examples.
package spin

import (
	"fmt"
	"io"
	"os"
	"time"
)

var standard *Spinner = New(StyleLine, time.Millisecond*100)

func GoStart() {
	standard.GoStart()
}
func Start() {
	standard.Start()
}
func Stop() {
	standard.Stop()
}

var out io.Writer = os.Stdout

var (
	// StyleLine is a simple example of the kinds of styles you could pass into the New... functions.
	StyleLine = []string{"|", "/", "-", "\\"}

	// StyleSteps is another example of the kinds of styles you could pass into the New... functions.
	StyleSteps = []string{"▁", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃"}

	// StyleShutter is another example of the kinds of styles you could pass into the New... functions.
	StyleShutter = []string{"▉", "▊", "▋", "▌", "▍", "▎", "▏", "▎", "▍", "▌", "▋", "▊", "▉"}
)

type Spinner struct {
	style  []string
	delay  time.Duration
	prefix string
	suffix string
	stop   chan struct{}
}

// New creates a spinner which you can start and stop.
func New(style []string, delay time.Duration) *Spinner {
	return &Spinner{
		style: style,
		delay: delay,
		stop:  make(chan struct{}),
	}
}

// NewWithPadding creates a spinner which you can start and stop.
// Allows a prefix and suffix to be printed along with the specified style.
func NewWithPadding(style []string, delay time.Duration, prefix, suffix string) *Spinner {
	return &Spinner{
		style:  style,
		delay:  delay,
		prefix: prefix,
		suffix: suffix,
		stop:   make(chan struct{}),
	}
}

// GoStart begins the spinner on a fresh goroutine.
func (self *Spinner) GoStart() {
	go self.Start()
}

// Start begins the spinner on the current goroutine (hopefully you've got another goroutine that can call Stop...).
func (self *Spinner) Start() {
	for {
		select {
		case <-self.stop:
			fmt.Fprint(out, "\r") // erase any residual spinner markings
			return
		default:
			self.spinCycle()
		}
	}
}
func (self *Spinner) spinCycle() {
	for _, symbol := range self.style {
		fmt.Fprintf(out, "\r%s%s%s", self.prefix, symbol, self.suffix)
		time.Sleep(self.delay)
	}
}

// Stop sends a signal to stop the spinner.
func (self *Spinner) Stop() {
	self.stop <- struct{}{}
}
