// Here several spinners are demonstrated. Try adding your own []string patterns.,,
package main

import (
	"fmt"
	"time"

	"github.com/mdwhatcott/spin"
)

func main() {
	fmt.Println()
	Show("spin.StyleLine        ", spin.StyleLine)
	Show("spin.StyleSteps       ", spin.StyleSteps)
	Show("spin.StyleShutter     ", spin.StyleShutter)
	fmt.Println("\nDone.")
}

func Show(title string, style []string) {
	spinner := spin.NewWithPadding(style, time.Millisecond*100, title, "    "+fmt.Sprint(style))
	spinner.GoStart()
	time.Sleep(time.Second * 3)
	spinner.Stop()
	fmt.Println("\n")
}
