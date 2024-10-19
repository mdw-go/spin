// Here several spinners are demonstrated. Try adding your own []string patterns.,,
package main

import (
	"fmt"
	"time"

	"github.com/mdw-go/spin"
)

func main() {
	ShowDefault()
	ShowStyles()
	fmt.Println("\nDone.")
}

func ShowDefault() {
	fmt.Println("Default spinner: ")
	spin.GoStart()
	time.Sleep(time.Second * 1)
	spin.Stop()
}

func ShowStyles() {
	styles := map[string]spin.Style{
		"spin.StylePops        ": spin.StylePops,
		"spin.StyleBrackets    ": spin.StyleBrackets,
		"spin.StyleLine        ": spin.StyleLine,
		"spin.StyleSteps       ": spin.StyleSteps,
		"spin.StyleShutter     ": spin.StyleShutter,
		"spin.StyleNumbers     ": spin.StyleNumbers,
		"spin.StyleAlphabet    ": spin.StyleAlphabet,
	}
	for title, style := range styles {
		Show(title, style)
	}
}

func Show(title string, style spin.Style) {
	fmt.Println()
	fmt.Println()
	spinner := spin.New(
		spin.Options.Style(style),
		spin.Options.Prefix(title),
		spin.Options.Suffix("    "+fmt.Sprint(style)),
	)
	go spinner.Start()
	time.Sleep(time.Second * 1)
	spinner.Stop()
}
