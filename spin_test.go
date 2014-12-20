package spin

import (
	"bytes"
	"strings"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpin(t *testing.T) {
	Convey("Subject: inspecting raw spin output", t, func() {
		output := &carriageReturnAbsorber{Buffer: bytes.NewBufferString("")}
		out = output

		Convey("As a result of starting a spinner", func(c C) {
			spin(New(StyleLine, time.Millisecond))

			Convey("The provided pattern should be repeatedly written to the output", func() {
				expected := strings.Join(StyleLine, "")
				So(output.String(), ShouldStartWith, expected+expected)
			})

			Convey("Each write to the console should start with a carriage return", func() {
				So(output.CarriageReturns, ShouldEqual, output.TotalWrites)
			})
		})

		Convey("When a prefix and suffix are provided", func(c C) {
			spin(NewWithPadding(StyleLine, time.Millisecond, ">> ", " <<"))

			Convey("The prefix and suffix should appear before and after each character in the pattern", func() {
				expected := ">> | <<" +
					">> / <<" +
					">> - <<" +
					">> \\ <<"
				So(output.String(), ShouldStartWith, expected+expected)
			})

			Convey("Each write to the console should start with a carriage return", func() {
				So(output.CarriageReturns, ShouldEqual, output.TotalWrites)
			})
		})
	})
}

func spin(spinner *Spinner) {
	spinner.Start()
	time.Sleep(time.Millisecond * 10)
	spinner.Stop()
}

///////////////////////////////////////////////////////////////////////////////

// carriageReturnAbsorber is an io.Writer that stores an internal buffer
// allowing inspection of what was written. It substitutes for os.Stdout
// but throws out leading carriage returns to make the tests easier to read.
type carriageReturnAbsorber struct {
	*bytes.Buffer
	CarriageReturns int
	TotalWrites     int
}

func (self *carriageReturnAbsorber) Write(value []byte) (int, error) {
	self.TotalWrites++

	if value[0] == '\r' && len(value) > 1 {
		self.CarriageReturns++
		return self.Buffer.Write(value[1:])
	} else {
		return self.Buffer.Write(value)
	}
}
