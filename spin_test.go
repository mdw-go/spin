package spin

import (
	"bytes"
	"sync/atomic"
	"testing"
	"time"

	"github.com/mdw-go/testing/should"
)

func TestSpinFixture(t *testing.T) {
	should.Run(&SpinFixture{T: should.New(t)}, should.Options.UnitTests())
}

type SpinFixture struct {
	*should.T

	absorber *carriageReturnAbsorber
	output   *output
}

func (this *SpinFixture) Setup() {
	this.absorber = newCarriageReturnAbsorber()
	this.output = &output{out: this.absorber}
}

func (this *SpinFixture) spin(spinner *Spinner) {
	spinner.out = this.output
	go spinner.Start()
	time.Sleep(time.Millisecond)
	spinner.Stop()
}

func (this *SpinFixture) TestSpinner() {
	this.spin(New(Options.Style(StyleLine), Options.Delay(time.Nanosecond)))

	expected := StyleLine

	this.Println("- The provided pattern should be written repeatedly to the output.")
	this.So(this.absorber.String(), should.StartWith, string(expected+expected))

	this.Println("- Each write should start with a carriage return.")
	this.So(this.absorber.CarriageReturns.Load(), should.Equal, this.absorber.TotalWrites.Load())
}

func (this *SpinFixture) TestSpinner_PrefixSuffix() {
	this.spin(New(Options.Style(StyleLine), Options.Delay(time.Nanosecond), Options.Prefix(">> "), Options.Suffix(" <<")))

	expected := ">> | <<" +
		">> / <<" +
		">> - <<" +
		">> \\ <<"
	this.Println("- The prefix and suffix should surround each write.")
	this.So(this.absorber.String(), should.StartWith, string(expected+expected))
}

///////////////////////////////////////////////////////////////////////////////

// carriageReturnAbsorber is an io.Writer that stores an internal buffer
// allowing inspection of what was written. It substitutes for os.Stdout
// but throws out leading carriage returns to make the tests easier to read.
type carriageReturnAbsorber struct {
	*bytes.Buffer
	CarriageReturns *atomic.Int32
	TotalWrites     *atomic.Int32
}

func newCarriageReturnAbsorber() *carriageReturnAbsorber {
	return &carriageReturnAbsorber{
		Buffer:          bytes.NewBufferString(""),
		CarriageReturns: new(atomic.Int32),
		TotalWrites:     new(atomic.Int32),
	}
}

func (self *carriageReturnAbsorber) Write(value []byte) (int, error) {
	self.TotalWrites.Add(1)

	if value[0] == '\r' && len(value) > 1 {
		self.CarriageReturns.Add(1)
		return self.Buffer.Write(value[1:])
	} else if value[0] == '\r' {
		self.CarriageReturns.Add(1)
		return 0, nil
	} else {
		return self.Buffer.Write(value)
	}
}

////////////////////////////////////////////////////////////////////////////
