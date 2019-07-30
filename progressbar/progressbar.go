package progressbar

import(
	"fmt"
	"reflect"
	"errors"
)

var White 	= 47
var Red   	= 41
var Green 	= 42
var Yellow	= 43
var Blue	= 44
var Magenta = 45
var Cyan  	= 46

type pbar struct {
	currLen int
	totalLen int
	isColor bool
	color int
	width int
	fillString string
	bar string
	LOW int
	MEDIUM int
	HIGH int
	FULL int
}

// @param length int
// @param length Slice
// @param pbarWidth int 20
// @param inColor bool false
func NewPbar(args ...interface{}) (*pbar, error) {
	var err error
	if 1 > len(args) {
        err = errors.New("Not enough parameters.")
        return nil, err
    }

	pb := new(pbar)
	pb.width = 20
	pb.isColor = false
	pb.fillString = " "

	pb.currLen = 0
	pb.LOW = Red
	pb.MEDIUM = Yellow
	pb.HIGH = Green
	pb.FULL = White

	// Validate args passed in, most importantly the first: will accept 
	//   either an int or a slice to determine the size of each chunk of the bar.
	for i,p := range args {
		switch i {
		case 0:
			switch reflect.TypeOf(p).Kind() {
		    case reflect.Slice:
		    	x := reflect.ValueOf(p).Len()
		        pb.totalLen = x
		    case reflect.Int:
		    	pb.totalLen = int(reflect.ValueOf(p).Int())
		    default:
		    	err = errors.New("First argument must be an Int or a Slice to determine length.")
		    	return nil, err
		    }
		case 1: // Determines the width of the progressbar
			param, ok := p.(int)
			if !ok {
                err = errors.New("2nd parameter not type int.")
                return nil, err
            }
            if param < 1 {
            	err = errors.New("Width must be at least 1.")
            	return nil, err
            }
            pb.width = param
        case 2: // Is in color
        	paramBool, ok := p.(bool)
        	paramInt, ok2 := p.(int)
        	if !(ok || ok2) {
        		err = errors.New("3rd parameter must be bool or int (0 = false).")
        		return nil, err
        	}
        	if !ok {
        		fmt.Println("hmm")
        		pb.isColor = (paramInt != 0)
        	} else {
        		pb.isColor = paramBool
        	}
		}
	}

	pb.prep()
	return pb, nil
}

func (pb * pbar) toggleColor(inColor bool) {
	pb.isColor = inColor
}

func (pb * pbar) setFillString(str string) {
	pb.fillString = str
}

var BarItem = "\033[47m"

var colors = []int{Red, Yellow, Green, White}

func (pb * pbar) prep() {
			// \033[0E 	: beginning of line (?)
			// [		: just prints out '[' string literal
			// \0337 	: store cursor position
			// \033[52G	: move cursor 52 spaces right
			// ]		: just prints out ']' string literal
			// \033[2G	: load stored cursor position
	fmt.Printf("\033[0E[\0337\033[%dG]\033[2G", pb.width+2)
	pb.bar = ""
	if !pb.isColor {
		pb.color = White
	}
}

//TODO: Change params to take slices rather than color ints
func (pb * pbar) changeColor() {
	if pb.currLen <= pb.totalLen/4 {
		pb.color = pb.LOW
	} else if pb.currLen > pb.totalLen/4 && pb.currLen <= 2*pb.totalLen/3 {
		pb.color = pb.MEDIUM
	} else if pb.currLen > 2*pb.totalLen/3 && pb.currLen < pb.totalLen {
		pb.color = pb.HIGH
	} else {
		pb.color = pb.FULL
	}

	if pb.fillString != " " {
		pb.color -= 10
	}
	
}

func (pb * pbar) increaseBar() {
	threshold := (float32(pb.totalLen) / float32(pb.width))+1
	pb.currLen+=1
	pb.bar = ""
	for i := 0; i < pb.currLen; i+=int(threshold) {
		pb.bar+=pb.fillString
	}
	pb.drawBar()
	if pb.currLen == pb.totalLen {
		fmt.Printf("\033[0m\n")
	}
}

func (pb *pbar) drawBar() {
	if pb.isColor {
		pb.changeColor()
	}
	fmt.Printf("\033[2G\033[%dm%s", pb.color, pb.bar)
}