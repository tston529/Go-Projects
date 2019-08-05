package progressbar

import (
	"errors"
	"fmt"
	"reflect"
)
import "github.com/k0kubun/go-ansi"

// NewPbar is a constructor for a progressbar object.
// @param length int
// @param length Slice
// @param pbarWidth int 20
// @param inColor bool false
func NewPbar(args ...interface{}) (*Pbar, error) {
	var err error
	if 1 > len(args) {
		err = errors.New("not enough parameters")
		return nil, err
	}

	if len(colors) == 0 {
		colors = map[string]int{
			"white":   47,
			"red":     41,
			"green":   42,
			"yellow":  43,
			"blue":    44,
			"magenta": 45,
			"cyan":    46,
		}
	}

	pb := new(Pbar)
	pb.forever = make(chan struct{})

	pb.incrementAmount = 1
	pb.unprepped = true
	pb.width = 20
	pb.isColor = false
	pb.fgString = "="
	pb.bgString = "_"
	pb.printNumbers = 0
	pb.currentAmount = 0
	pb.LOW = colors["red"]
	pb.MEDIUM = colors["yellow"]
	pb.HIGH = colors["green"]
	pb.FULL = colors["white"]

	// Validate args passed in, most importantly the first: will accept
	//   either an int or a slice to determine the size of each chunk of the bar.
	for i, p := range args {
		switch i {
		case 0:
			switch reflect.TypeOf(p).Kind() {
			case reflect.Slice:
				x := reflect.ValueOf(p).Len()
				pb.totalAmount = x
			case reflect.Int:
				pb.totalAmount = int(reflect.ValueOf(p).Int())
			default:
				err = errors.New("first argument must be an Int or a Slice to determine length")
				return nil, err
			}
		case 1: // Determines the width of the progressbar
			param, ok := p.(int)
			if !ok {
				err = errors.New("2nd parameter not type int")
				return nil, err
			}
			if param < 1 {
				err = errors.New("width must be at least 1")
				return nil, err
			}
			pb.width = param
		}
	}
	pb.blockThreshold = int((float32(pb.totalAmount) / float32(pb.width)) + 1)
	return pb, nil
}

// ToggleColor lets the user set whether the bar is color or b/w
func (pb *Pbar) ToggleColor(arg interface{}) {
	switch reflect.TypeOf(arg).Kind() {
	case reflect.Slice:
		col := reflect.ValueOf(arg)
		l := col.Len()

		if l < 0 || l > 4 {
			fmt.Println("Error: Must provide a number of colors between 0 - 4.\nSwitching to b/w...")
			pb.isColor = false
			return
		}

		for i := 0; i < col.Len(); i++ {
			colStr := col.Index(i).String()
			if _, ok := colors[colStr]; !ok {
				fmt.Printf("'%s' is not a valid color. Switching to b/w...\n\n", colStr)
				pb.isColor = false
				return
			}
		}

		if l == 4 {
			pb.LOW = colors[(col.Index(0)).String()]
			pb.MEDIUM = colors[(col.Index(1)).String()]
			pb.HIGH = colors[(col.Index(2)).String()]
			pb.FULL = colors[(col.Index(3)).String()]
		} else if l == 3 {
			pb.LOW = colors[(col.Index(0)).String()]
			pb.MEDIUM = colors[(col.Index(1)).String()]
			pb.HIGH = colors[(col.Index(2)).String()]
			pb.FULL = colors[(col.Index(2)).String()]
		} else if l == 2 {
			pb.LOW = colors[(col.Index(0)).String()]
			pb.MEDIUM = colors[(col.Index(0)).String()]
			pb.HIGH = colors[(col.Index(1)).String()]
			pb.FULL = colors[(col.Index(1)).String()]
		} else if l == 1 {
			pb.LOW = colors[(col.Index(0)).String()]
			pb.MEDIUM = colors[(col.Index(0)).String()]
			pb.HIGH = colors[(col.Index(0)).String()]
			pb.FULL = colors[(col.Index(0)).String()]
		} else {
			pb.isColor = false
			return
		}
		pb.isColor = true
	case reflect.String:
		col := reflect.ValueOf(arg)
		if _, ok := colors[col.String()]; ok {
			pb.LOW = colors[col.String()]
			pb.MEDIUM = colors[col.String()]
			pb.HIGH = colors[col.String()]
			pb.FULL = colors[col.String()]
			pb.isColor = true
		}
	}
}

// SetGraphics lets the user change the fg/bg text
func (pb *Pbar) SetGraphics(args ...string) {
	if len(args) == 0 {
		pb.fgString = " "
		pb.bgString = " "
	} else if len(args) == 1 {
		pb.fgString = args[0]
	} else if len(args) == 2 {
		pb.fgString = args[0]
		pb.bgString = args[1]
	} else {
		fmt.Println("Too many arguments. Only need two.")
	}
	ansi.Printf("\033[2G")
}

// SetWidth allows the user to change how many characters wide the bar is drawn
func (pb *Pbar) SetWidth(arg int) {
	if arg < 5 {
		arg = 5
	}
	pb.width = arg
}

// SetIncrementAmt allows the user to change the incrementation amount (default is 1)
func (pb *Pbar) SetIncrementAmt(arg int) {
	if arg <= 0 {
		arg = 1
	}
	pb.incrementAmount = arg
}

// SetPrintNumbers gives the user the option to display progress numerically
//   in addition to the regular bar. Options allow for a percent or a fraction.
func (pb *Pbar) SetPrintNumbers(arg string) {
	if arg == "percent" || arg == "%" {
		pb.printNumbers = 2
	} else if arg == "ratio" || arg == "numeric" || arg == "/" || arg == "fraction" {
		pb.printNumbers = 1
	} else {
		pb.printNumbers = 0
	}
}

// IncreaseBar increments its internal counter and recalculates the
//   size of the bar that is filled in.
func (pb *Pbar) IncreaseBar() {

	// Run initialization stuff if not already.
	//   Need to do this now because of extra stuff user
	//   may change before running their loop.
	if pb.unprepped {
		pb.prep()
		// fmt.Println(pb.fgStringLen)
	}

	pb.currentAmount += pb.incrementAmount
	if pb.currentAmount/pb.blockThreshold >= pb.barLen {
		pb.filledBar += pb.fgString
		if pb.isColor {
			pb.changeColor()
		}
		pb.drawBar()
		pb.barLen += pb.fgStringLen
	}
	/* prevBar := pb.filledBar
	pb.filledBar = ""
	var i int
	for i = 0; i < pb.currentAmount; i += pb.blockThreshold {
		pb.filledBar += pb.fgString
	}
	if len(pb.filledBar) > len(prevBar) {
		pb.drawBar()
	} */
	if pb.printNumbers > 0 {
		ansi.Printf("\u001b[1B")
		pb.printValues()
	}
	if pb.currentAmount >= pb.totalAmount {
		ansi.Printf("\033[0m\n\n")
		ansi.CursorShow()
	}
}
