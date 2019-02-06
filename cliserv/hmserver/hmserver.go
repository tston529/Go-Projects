package main

import (	
	"os"
	"fmt"
	"net"
	"time"
	"bufio"
	"strconv"
	"strings"
	"math/rand"
)

var clientConns []net.Conn
var incomplete_word []string
var guesses []string
var hangman = 7
var badGuesses = 0
var wonGame = false

var RED = "\033[91m"
var GREEN = "\033[92m"
var BLUE = "\033[36m"
var YELLOW = "\033[33m"
var NO_COLOR = "\033[0m"

func check(err error, message string) {
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", message)
}

type ClientJob struct {
	init bool
	guess string
	conn net.Conn
}

func generateResponses(clientJobs chan ClientJob, the_word string) {
	for badGuesses < hangman && !wonGame {
		// Wait for the next job to come off the queue.
		clientJob := <-clientJobs

		g := strings.TrimSpace(clientJob.guess)
		if g == ""{
			continue
		}
		if clientJob.init {
			clientJob.conn.Write([]byte(clientJob.guess + "\n"))
		} else {
			_ = verifyGuess(g, the_word, incomplete_word)
			parseGuesses(g, the_word)
			if strings.Join(incomplete_word, "") == the_word {
				wonGame = true
			}
			for _, c := range clientConns {
				// Write badGuesses
				var bg strings.Builder
				bg.WriteString(strconv.Itoa(badGuesses))
				c.Write([]byte(bg.String() + "\n"))
				time.Sleep(10 * time.Millisecond)

				// Write incomplete_word
				c.Write([]byte(strings.Join(incomplete_word, "") + "\n"))
				time.Sleep(10 * time.Millisecond)

				// Write guesses
				c.Write([]byte(strings.Join(guesses, ", ") + "\n"))
				time.Sleep(10 * time.Millisecond)

				// Write win status
				if wonGame {
					c.Write([]byte("WIN\n"))
				} else {
					c.Write([]byte("CONT\n"))
				}
			}
			continue
		}

	}

	endGame(wonGame, the_word)
	os.Exit(1)
}

func genWord(dictionary []string)(chosen_word string) {
	if dictionary==nil{
		dictionary = []string{"computer", "guitar", "clock", "bookcase", "cactus", "succulent"}
	}
	chosen_word = dictionary[rand.Intn(len(dictionary))]
	return 
}

// Perform some misc. prep work
func initialize() string {
	rand.Seed(time.Now().UTC().UnixNano())
	dictionary := openFile()
	return genWord(dictionary)
}

// Open a text file dictionary. Return it if available, otherwise return nil
func openFile()( lines []string) {
	file, err := os.Open("./eng.txt")
	if err != nil {
		fmt.Print(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}

func verifyGuess(guess string, the_word string, incomplete_word []string) bool {
	found := false
	for x := 0; x < len(the_word); x++ {
		if string(the_word[x]) == guess {
			incomplete_word[x] = guess
			found = true
		}
	}
	return found
}

func parseGuesses(guess string, the_word string) {
	var rslt strings.Builder
	if(strings.Contains(the_word, guess)){
		rslt.WriteString(GREEN)
	} else {
		rslt.WriteString(RED)
		badGuesses++
		fmt.Println("badGuesses: %d", badGuesses)
	}
	rslt.WriteString(guess)
	rslt.WriteString(NO_COLOR)
	guesses = append(guesses, rslt.String())
}

func endGame(won_game bool, the_word string) {
	fmt.Println()
	if(won_game) {
		fmt.Println("Congrats!")
	} else {
		fmt.Printf("RIP: The word was \"%s\"\n", the_word)
	}
}

func main() {
	
	the_word := strings.TrimRight(initialize(), "\n")

	for range the_word {
		incomplete_word = append(incomplete_word, "_")
	}

	clientJobs := make(chan ClientJob)
	go generateResponses(clientJobs, the_word)
	
	ln, err := net.Listen("tcp", ":15325")
	check(err, "Server is ready.")

	for {
		conn, err := ln.Accept()
		check(err, "Accepted connection.")

		clientJobs <- ClientJob{true, the_word, conn}

		var bg strings.Builder
		bg.WriteString(strconv.Itoa(badGuesses))

		clientJobs <- ClientJob{true, bg.String(), conn}

		inc := strings.Join(incomplete_word, "")
		clientJobs <- ClientJob{true, inc, conn}

		clientConns = append(clientConns, conn)
		
		fmt.Print(clientConns)


		go func() {
			buf := bufio.NewReader(conn)

			for {
				guess, err := buf.ReadString('\n')

				if err != nil {
					fmt.Printf("Client disconnected.\n")
					break
				}

				clientJobs <- ClientJob{false, guess, conn}
			}
		}()
	}
}
