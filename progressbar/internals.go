package progressbar

import "fmt"
import "context"

var colors map[string]int

// pbar contains the data for the progress bar.
type pbar struct {
    forever chan struct{}
    unprepped bool

    incrementAmount int
    currentAmount int
    totalAmount int
    blockSize int

    printNumbers int
    isColor bool
    color int
    width int

    // Typically a single character each, the object
    //   to be drawn representing a single block of the bar
    fgString string
    bgString string

    // A longer, chained-together version of the fg/bgString counterparts
    //   representing the current fullness/emptiness of the bar, respectively.
    filledBar string
    bgBar string

    // these hold ansi terminal data for colors, to be used
    //   in various points based on the progressbar's fullness.
    LOW int
    MEDIUM int
    HIGH int
    FULL int
}

// prep finalizes the progressbar by drawing the background
//  and contextually spinning up a daemon to switch colors
//  in the background.
func (pb *pbar) prep() {
    // \033[0E     : beginning of line (?)
    // [        : just prints out '[' string literal
    // \0337     : store cursor position
    // \033[52G    : move cursor 52 spaces right
    // ]        : just prints out ']' string literal
    // \033[2G    : load stored cursor position
    pb.filledBar = ""
    if !pb.isColor {
        pb.color = colors["white"]
        // If b/w, make the text visible
        //   (otherwise background would print white)
        if pb.fgString != " " {
            pb.color -= 10
        }
    }

    // Create the background, then print it
    for i := 0; i < pb.width; i+=1 {
        pb.bgBar+=pb.bgString
    }
    fmt.Printf("\033[0E[\0337%s]\033[2G", pb.bgBar)

    // If the user has opted to print a numerical value,
    //   print it on the next line
    if pb.printNumbers > 0 {
        fmt.Println()
        pb.printValues()
    }
    // Pbar is now prepped and no longer has to run this function.
    pb.unprepped = false

    // If user has chosen to colorize the progressbar,
    //   spin up the color switcher daemon.
    if pb.isColor {
        ctx, cancel := context.WithCancel(context.Background())
        go pb.colorDaemon(ctx)
        go func() {
            <-pb.forever
            for {
                if pb.currentAmount >= pb.totalAmount {
                    cancel()
                }
            }
        }()
    }
}

// colorDaemon, when invoked, is constantly spinning in
//   the background to change the color of the bar.
func (pb *pbar) colorDaemon(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():  // if cancel() execute
            pb.forever <- struct{}{}
            return
        default:
            pb.changeColor()
        }
    }
}

// cjangeColor changes the current color to be drawn based on current progress
func (pb *pbar) changeColor() {
    if pb.fgString != " " {
        pb.color = -10
    } else {
        pb.color = 0
    }
    if pb.currentAmount <= pb.totalAmount/4 {
        pb.color += pb.LOW
    } else if pb.currentAmount > pb.totalAmount/4 && pb.currentAmount <= 2*pb.totalAmount/3 {
        pb.color += pb.MEDIUM
    } else if pb.currentAmount > 2*pb.totalAmount/3 && pb.currentAmount < pb.totalAmount {
        pb.color += pb.HIGH
    } else {
        pb.color += pb.FULL
    }

}

// drawBar writes the current state of the progressbar to the terminal
//   along with the numerical output, if selected.
func (pb *pbar) drawBar() {
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