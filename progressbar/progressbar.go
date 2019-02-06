package main

import(
	"fmt"
	"time"
)

type Range struct {
	gScalar int
	gDivisor int
	lScalar int
	lDivisor int
}

var color int
var bar string

var White = 47
var Green = 42
var Yellow= 43
var Red   = 41
var TotalLen = 20
var BarItem = "\033[47m"

var sleepTime = 50
var colors = []int{Red, Yellow, Green, White}
var ranges = []Range{
	Range{
		gScalar:0,
		gDivisor:1,
		lScalar:1,
		lDivisor:4 }, 
	Range{
		gScalar:1,
		gDivisor:4,
		lScalar:2,
		lDivisor:3 }, 
	Range{
		gScalar:2,
		gDivisor:3,
		lScalar:1,
		lDivisor:1 } }

func main() {
	prep()
	increaseBar(TotalLen, sleepTime)
}

func prep() {
			// \033[0E 	: beginning of line (?)
			// [		: just prints out '[' string literal
			// \0337 	: store cursor position
			// \033[52G	: move cursor 52 spaces right
			// ]		: just prints out ']' string literal
			// \033[2G	: load stored cursor position
	fmt.Printf("\033[0E[\0337\033[%dG]\033[2G", TotalLen+2)
	bar = ""
}

//TODO: Change params to take slices rather than color ints
func changeColor(currentAmt int, LOW int, MEDIUM int, HIGH int, FULL int) {
	if currentAmt <= TotalLen/4 {
		color = LOW
	} else if currentAmt > TotalLen/4 && currentAmt <= 2*TotalLen/3 {
		color = MEDIUM
	} else if currentAmt > 2*TotalLen/3 && currentAmt < TotalLen {
		color = HIGH
	} else {
		color = FULL
	}
	
}

func compareRange(currentAmt int, gScalar int, gDivisor int, lScalar int, lDivisor int) bool {
	return (currentAmt > gScalar*TotalLen/gDivisor && currentAmt <= lScalar*TotalLen/lDivisor);
}

func increaseBar(realNum int, timeMS int) {
	var barItemLen float32 
	barItemLen = float32(TotalLen)/float32(realNum)

	var i float32
	for i = 0; i < float32(TotalLen); i+=barItemLen {
		bar+=" "
		drawBar(int(i))
		time.Sleep(time.Duration(timeMS)*time.Millisecond)
	}
	
	
	fmt.Printf("\033[0m\n")
}

func drawBar(realNum int) {
	changeColor(int(realNum), Red, Yellow, Green, White)
	fmt.Printf("\033[2G\033[%dm%s", color, bar)
}