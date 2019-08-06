package progressbar

import (
	"fmt"
	"strconv"

	"github.com/k0kubun/go-ansi"
)

var colors map[string]int

// Pbar contains the data for the progress bar.
type Pbar struct {
	// The other option would be to have the user
	//   call the prep function manually, and I don't
	//   trust the end user to do this.
	unprepped bool

	// Typical value will be 1, but in case user opts to
	//   have a loop with a different increment amt
	incrementAmount int

	currentAmount int
	totalAmount   int

	// (totalAmount / pb.width)
	//   bar adds another instance of the fg element when
	//   another round of the threshold has been reached.
	blockThreshold int

	// Current length of the drawn bar.
	//   Increments by the drawable length of the fg string
	barLen int

	// Flag: Dictates if and how numerical progress will
	//   be printed under the progress bar.
	printNumbers int
	// Flag: Dictates if bar will be color or b/w
	isColor bool

	// Current color to be printed; number is ansi code
	color int

	// Width of the bar in characters, not including endcaps
	width int

	// Typically a single character each, the object
	//   to be drawn representing a single block of the bar
	fgString string
	bgString string
	// Typical value is going to be 1, but in the case a user
	//   specifies a foreground element larger than one,
	//   this will keep track of that size to only print
	//   out as few as necessary.
	fgStringLen int

	// A longer, chained-together version of the fg/bgString counterparts
	//   representing the current fullness/emptiness of the bar, respectively.
	filledBar string
	bgBar     string

	// these hold ansi terminal data for colors, to be used
	//   in various points based on the progressbar's fullness.
	LOW    int
	MEDIUM int
	HIGH   int
	FULL   int
}

// prep finalizes the progressbar by drawing the background
//  and preparing space for the bar and numeric progress, if necessary.
func (pb *Pbar) prep() {
	ansi.CursorHide()

	pb.filledBar = ""
	pb.barLen = 0
	pb.fgStringLen = parseBlockSize(pb.fgString)
	if !pb.isColor {
		pb.color = colors["white"]
		// If b/w, make the text visible
		//   (otherwise background would print white)
		if pb.fgString != " " {
			pb.color -= 10
		}
	}

	// Create the background, then print it
	for i := 0; i < pb.width; i++ {
		pb.bgBar += pb.bgString
	}

	// \033[0E     : beginning of line (?)
	// [        : just prints out '[' string literal
	// \0337     : store cursor position
	// \033[52G    : move cursor 52 spaces right
	// ]        : just prints out ']' string literal
	// \033[2G    : load stored cursor position
	ansi.Printf("\033[0E[\0337%s]\033[2G", pb.bgBar)

	// If the user has opted to print a numerical value,
	//   print it on the next line
	if pb.printNumbers > 0 {
		fmt.Println()
		pb.printValues()
	}
	// Pbar is now prepped and no longer has to run this function.
	pb.unprepped = false

}

// changeColor changes the current color to be drawn based on current progress
func (pb *Pbar) changeColor() {
	if pb.fgString != " " {
		pb.color = -10
	} else {
		pb.color = 0
	}
	if pb.currentAmount <= pb.totalAmount/4 {
		pb.color += pb.LOW
	} else if pb.currentAmount >= pb.totalAmount/4 && pb.currentAmount <= 2*pb.totalAmount/3 {
		pb.color += pb.MEDIUM
	} else if pb.currentAmount >= 2*pb.totalAmount/3 && pb.currentAmount <= pb.totalAmount {
		pb.color += pb.HIGH
	} else {
		pb.color += pb.FULL
	}

}

// drawBar writes the current state of the progressbar to the terminal
//   along with the numerical output, if selected.
func (pb *Pbar) drawBar() {
	ansi.Printf("\033[2G\033[%dm%s", pb.color, pb.filledBar)
}

// printValues prints the numerical values of progress, either in fractional
//   form or in a percentage, based on user specification.
func (pb *Pbar) printValues() {
	if pb.printNumbers == 1 {
		ansi.Printf("\033[2G\033[0m%d / %d\u001b[1A", pb.currentAmount, pb.totalAmount)
	} else if pb.printNumbers == 2 {
		ansi.Printf("\033[2G\033[0m%d%s\u001b[1A", int(100*float32(pb.currentAmount)/float32(pb.totalAmount)), "%")
	}
}

// parseBlockSize takes in a string and returns how many characters
//   will actually be displayed.  This is a very naive implementation,
//   as the only use cases in which it will work are when the final
//   x characters in the string are the "displayable" characters,
//   and where the ansi escape sequence(s), if any, end in 'm'
//   (the sequences which change font style: color, bold, italic, etc.)
func parseBlockSize(str string) int {
	visibleStrLen := 1
	i := len(str) - 1
	for {
		if i == 0 {
			break
		}
		if str[i] == byte('m') {
			if _, err := strconv.Atoi(string(str[i-1])); err == nil || i == 0 {
				visibleStrLen--
				break
			}
		}
		i--
		visibleStrLen++
	}
	// fmt.Println(len(str))
	return visibleStrLen
}
