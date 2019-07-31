package progressbar

import "fmt"

var colors map[string]int

type pbar struct {
	unprepped bool

	incrementAmount int
	currentAmount int
	totalAmount int
	blockSize int

	printNumbers int
	isColor bool
	color int
	width int

	fgString string
	bgString string

	filledBar string
	bgBar string

	LOW int
	MEDIUM int
	HIGH int
	FULL int
}

func (pb *pbar) prep() {
			// \033[0E 	: beginning of line (?)
			// [		: just prints out '[' string literal
			// \0337 	: store cursor position
			// \033[52G	: move cursor 52 spaces right
			// ]		: just prints out ']' string literal
			// \033[2G	: load stored cursor position
	pb.filledBar = ""
	if !pb.isColor {
		pb.color = colors["white"]
		// If b/w, make the text visible (otherwise background would print white)
		if pb.fgString != " " {
			pb.color -= 10
		}
	}
	for i := 0; i < pb.width; i+=1 {
		pb.bgBar+=pb.bgString
	}
	fmt.Printf("\033[0E[\0337%s]\033[2G", pb.bgBar)
	if pb.printNumbers > 0 {
		fmt.Println()
		pb.printValues()
	}
	fmt.Printf("")
	pb.unprepped = false
}

//TODO: Change params to take slices rather than color ints
func (pb *pbar) changeColor() {
	if pb.currentAmount <= pb.totalAmount/4 {
		pb.color = pb.LOW
	} else if pb.currentAmount > pb.totalAmount/4 && pb.currentAmount <= 2*pb.totalAmount/3 {
		pb.color = pb.MEDIUM
	} else if pb.currentAmount > 2*pb.totalAmount/3 && pb.currentAmount < pb.totalAmount {
		pb.color = pb.HIGH
	} else {
		pb.color = pb.FULL
	}

	if pb.fgString != " " {
		pb.color -= 10
	}
}


func (pb *pbar) drawBar() {
	if pb.isColor {
		pb.changeColor()
	}
	if pb.printNumbers > 0 {
		fmt.Printf("\u001b[1B")
		pb.printValues()
	}
	fmt.Printf("\033[2G\033[%dm%s", pb.color, pb.filledBar)
}

func (pb *pbar) printValues() {
	if pb.printNumbers == 1 {
		fmt.Printf("\033[2G\033[0m%d / %d\u001b[1A", pb.currentAmount, pb.totalAmount)
	} else if pb.printNumbers == 2 {
		fmt.Printf("\033[2G\033[0m%d%s\u001b[1A", int(100*float32(pb.currentAmount)/float32(pb.totalAmount)), "%")
	}
}