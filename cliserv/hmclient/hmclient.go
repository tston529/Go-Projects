package main

import (
	"os"
	"fmt"
	"net"
	"bufio"
	"strconv"
	"strings"
	"github.com/tston529/cls"
)

// Contains data about the string to use for the limb,
//   on which line to place it, and at which column.
type NextLimb struct {
	char string
	level int 
	column int
}

var badGuesses = 0
var wc = 0
var theWord string

var hangman = make([]NextLimb, 7)
var RED = "\033[91m"
var GREEN = "\033[92m"
var BLUE = "\033[36m"
var YELLOW = "\033[33m"
var NO_COLOR = "\033[0m"

var gallows = [] string { 
	"   /====\\   ",
	"   |    |   ",
	"   |        ",
	"   |        ",
	"   |        ",
	"   |        ",
	" __|_______ ",
	"/  +       \\",
	"+==========+"	}

func main() {
	var incomplete_word []string
	createLimbs()
	badGuesses = 0
	cls.CallClear()
	var conn = connectToServer()

	tw, _ := bufio.NewReader(conn).ReadString('\n')
    theWord = strings.TrimSuffix(tw, "\n")
	
	updateBoard(conn)
	incomplete_word = getIncompleteWord(conn)

	drawBoard()

	printSecretWord(incomplete_word)

	ch := make(chan struct{})
	go func() {
		fmt.Print("Enter your guess: ")
		for wc == 0 {
			reader := bufio.NewReader(os.Stdin)
		    text, _ := reader.ReadString('\n')
		    // send to socket
		    fmt.Fprintf(conn, text + "\n")
		}
		ch <- struct{}{}
	}()

	go func() {
		for wc == 0 { 
		    // read in input from stdin
		    // listen for reply
    		updateBoard(conn)
    		drawBoard()

		    incomplete_word = getIncompleteWord(conn)
		    printSecretWord(incomplete_word)
			
			message, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Printf("Previous guesses: [ ")
			fmt.Printf("%s", strings.TrimSuffix(message, "\n"))
			fmt.Printf("_ ]\n")

			wc = getWinCond(conn)

			if wc == 0 {
		    	fmt.Print("Enter your guess: ") 
			} else {
				break
			}
	  	}
	  	ch <- struct{}{}
	}()

	for wc == 0 {
        <-ch
    }

	// The core game loop
	end_game(wc==1, theWord)
}



// The position in the hangman[] array represents the 
//   order in which the limbs will be placed in the 
//   ascii gallows.
func createLimbs() {
	hangman[0] = (NextLimb{"o", 3, 9})
	hangman[1] = (NextLimb{"|", 4, 9})
	hangman[2] = (NextLimb{"/", 4, 8})
	hangman[3] = (NextLimb{"\\", 4, 10})
	hangman[4] = (NextLimb{"+", 5, 9})
	hangman[5] = (NextLimb{"/", 6, 8})
	hangman[6] = (NextLimb{"\\", 6, 10})
}

// Display the player's progress in finding the word
// e.g.:    p n _ _ m _ c _ c c _ s
func printSecretWord(incomplete_word []string) {
	for x := range incomplete_word {
		if incomplete_word[x] != "_" {
			fmt.Printf("%s%s%s ", YELLOW, incomplete_word[x], NO_COLOR)
		} else {
			fmt.Printf("%s ", incomplete_word[x])
		}
	}
	fmt.Printf("\n")
}

func updateBoard(conn net.Conn) {
	message, _ := bufio.NewReader(conn).ReadString('\n')
	bg, _ := strconv.Atoi(strings.TrimSuffix(message, "\n"))

    badGuesses = bg

    for i := 0; i < badGuesses; i++ {
    	limb := hangman[i]
    	gallows[limb.level-1] = replace_at_index(gallows[limb.level-1], limb.char, limb.column-1)
    }
}

func getIncompleteWord(conn net.Conn) []string {
	message, _ := bufio.NewReader(conn).ReadString('\n')
    iw := strings.TrimSuffix(message, "\n")
    var in_w []string
    for i := 0; i < len(iw); i++ {
    	in_w = append(in_w, string(iw[i]))
    }
    return in_w
}

func connectToServer() net.Conn {
	conn, _ := net.Dial("tcp", ":15325")
	return conn
}

func drawBoard() {
	for x := range gallows {
		fmt.Printf("\033[%d;0H%s", (x+2), gallows[x])
	}
	fmt.Println()
}

func getWinCond(conn net.Conn) int {
	message, _ := bufio.NewReader(conn).ReadString('\n')
	w := strings.TrimSpace(message)
	if w == "WIN" {
		return 1
	}
	if w == "CONT" {
		return 0
	}
	return -1
}

// Display feedback depending on whether the player won or lost
func end_game(won_game bool, theWord string) {
	fmt.Println()
	if won_game {
		fmt.Println("Congrats!")
	} else {
		fmt.Printf("RIP: The word was \"%s\"\n", theWord)
	}
}

// Helper function; for use in adding limbs to the hangman
func replace_at_index(str string, substring string, index int) string {
	return str[:index] + string(substring) + str[index+1:]
}
