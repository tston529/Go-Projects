package main

import (
	"fmt"
	//"math"
)

var BG_CYAN = 46
var BG_YELLOW = 43
var FG_CYAN = 96
var FG_YELLOW = 93

var unicodes = make([]string, 7)

// Board : the chessboard
var Board = make([]Cell, 64)

// AvailableMoves : List of available moves for any selected piece
var AvailableMoves = []int{}

// ChessPiece : a chess piece
type ChessPiece int

// Enumerated list of chess piece types
const (
	EMPTY ChessPiece = iota
	PAWN
	ROOK
	KNIGHT
	BISHOP
	QUEEN
	KING
)

// Color : color of chess piece
type Color int

// Enumerated list of chess colors
const (
	BLACK Color = -1
	NONE        = 0
	WHITE       = 1
)

// Position : relative to the piece's current position
type Position int

// Enumerated list of relative positions
const (
	FORWARD                = -8
	DIAGONAL_FORWARD_LEFT  = -9
	DIAGONAL_FORWARD_RIGHT = -7
	BACK                   = 8
	DIAGONAL_BACK_LEFT     = 7
	DIAGONAL_BACK_RIGHT    = 9
	LEFT                   = -1
	RIGHT                  = 1
	L1                     = -17
	L2                     = -15
	L3                     = -6
	L4                     = -10
	L5                     = 6
	L6                     = 10
	L7                     = 15
	L8                     = 17
)

// Piece : player-specific chess piece
type Piece struct {
	PieceType  ChessPiece
	PieceColor Color
	Moved      bool
}

func (pc *Piece) findMoves(board []Cell, currPos int) {
	switch pc.PieceType {
	case PAWN:
		if pc.Moved {
			if pc.checkEmptyCell(board, currPos, int(pc.PieceColor*FORWARD)) {
				AvailableMoves = append(AvailableMoves, int(pc.PieceColor*FORWARD))
			}
		} else {
			if pc.checkEmptyCell(board, currPos, int(pc.PieceColor*FORWARD)) {
				AvailableMoves = append(AvailableMoves, int(pc.PieceColor*FORWARD))
			}
			if pc.checkEmptyCell(board, currPos, int(pc.PieceColor*FORWARD+pc.PieceColor*FORWARD)) {
				AvailableMoves = append(AvailableMoves, int(pc.PieceColor*FORWARD+pc.PieceColor*FORWARD))
			}
		}
		break

	case ROOK:
		f := FORWARD
		b := BACK
		l := LEFT
		r := RIGHT
		for i := 0; i < 8; i++ {
			if pc.checkEmptyCell(board, currPos, f) {
				AvailableMoves = append(AvailableMoves, f)
				f += FORWARD
			}
			if pc.checkEmptyCell(board, currPos, b) {
				AvailableMoves = append(AvailableMoves, b)
				b += BACK
			}
			if pc.checkEmptyCell(board, currPos, l) {
				AvailableMoves = append(AvailableMoves, l)
				l += LEFT
			}
			if pc.checkEmptyCell(board, currPos, r) {
				AvailableMoves = append(AvailableMoves, r)
				r += RIGHT
			}

		}
		break
	case BISHOP:
		dfl := DIAGONAL_FORWARD_LEFT
		dfr := DIAGONAL_FORWARD_RIGHT
		dbl := DIAGONAL_BACK_LEFT
		dbr := DIAGONAL_BACK_RIGHT
		for i := 0; i < 8; i++ {
			if pc.checkEmptyCell(board, currPos, dfl) {
				AvailableMoves = append(AvailableMoves, dfl)
				dfl += DIAGONAL_FORWARD_LEFT
			}
			if pc.checkEmptyCell(board, currPos, dfr) {
				AvailableMoves = append(AvailableMoves, dfr)
				dfr += DIAGONAL_FORWARD_RIGHT
			}
			if pc.checkEmptyCell(board, currPos, dbl) {
				AvailableMoves = append(AvailableMoves, dbl)
				dbl += DIAGONAL_BACK_LEFT
			}
			if pc.checkEmptyCell(board, currPos, dbr) {
				AvailableMoves = append(AvailableMoves, dbr)
				dbr += DIAGONAL_BACK_RIGHT
			}

		}
		break
	case QUEEN:
		f := FORWARD
		b := BACK
		l := LEFT
		r := RIGHT
		dfl := DIAGONAL_FORWARD_LEFT
		dfr := DIAGONAL_FORWARD_RIGHT
		dbl := DIAGONAL_BACK_LEFT
		dbr := DIAGONAL_BACK_RIGHT
		for i := 0; i < 8; i++ {
			if pc.checkEmptyCell(board, currPos, f) {
				AvailableMoves = append(AvailableMoves, f)
				f += FORWARD
			}
			if pc.checkEmptyCell(board, currPos, b) {
				AvailableMoves = append(AvailableMoves, b)
				b += BACK
			}
			if pc.checkEmptyCell(board, currPos, l) {
				AvailableMoves = append(AvailableMoves, l)
				l += LEFT
			}
			if pc.checkEmptyCell(board, currPos, r) {
				AvailableMoves = append(AvailableMoves, r)
				r += RIGHT
			}
			if pc.checkEmptyCell(board, currPos, dfl) {
				AvailableMoves = append(AvailableMoves, dfl)
				dfl += DIAGONAL_FORWARD_LEFT
			}
			if pc.checkEmptyCell(board, currPos, dfr) {
				AvailableMoves = append(AvailableMoves, dfr)
				dfr += DIAGONAL_FORWARD_RIGHT
			}
			if pc.checkEmptyCell(board, currPos, dbl) {
				AvailableMoves = append(AvailableMoves, dbl)
				dbl += DIAGONAL_BACK_LEFT
			}
			if pc.checkEmptyCell(board, currPos, dbr) {
				AvailableMoves = append(AvailableMoves, dbr)
				dbr += DIAGONAL_BACK_RIGHT
			}

		}
		break
	case KING:
		f := FORWARD
		b := BACK
		l := LEFT
		r := RIGHT
		dfl := DIAGONAL_FORWARD_LEFT
		dfr := DIAGONAL_FORWARD_RIGHT
		dbl := DIAGONAL_BACK_LEFT
		dbr := DIAGONAL_BACK_RIGHT
		if pc.checkEmptyCell(board, currPos, f) {
			AvailableMoves = append(AvailableMoves, f)
		}
		if pc.checkEmptyCell(board, currPos, b) {
			AvailableMoves = append(AvailableMoves, b)
		}
		if pc.checkEmptyCell(board, currPos, l) {
			AvailableMoves = append(AvailableMoves, l)
		}
		if pc.checkEmptyCell(board, currPos, r) {
			AvailableMoves = append(AvailableMoves, r)
		}
		if pc.checkEmptyCell(board, currPos, dfl) {
			AvailableMoves = append(AvailableMoves, dfl)
		}
		if pc.checkEmptyCell(board, currPos, dfr) {
			AvailableMoves = append(AvailableMoves, dfr)
		}
		if pc.checkEmptyCell(board, currPos, dbl) {
			AvailableMoves = append(AvailableMoves, dbl)
		}
		if pc.checkEmptyCell(board, currPos, dbr) {
			AvailableMoves = append(AvailableMoves, dbr)
		}
		break
	case KNIGHT:
		if pc.checkEmptyCell(board, currPos-8, -8) {
			fmt.Println("OkL1")
			if pc.checkEmptyCell(board, currPos-8-8, -1) {
				AvailableMoves = append(AvailableMoves, L1)
			}
		}
		if pc.checkEmptyCell(board, currPos-8, -8) {
			fmt.Printf("OkL2")
			if pc.checkEmptyCell(board, currPos-8-8, 1) {
				AvailableMoves = append(AvailableMoves, L2)
			}
		}
		if pc.checkEmptyCell(board, currPos, L3) {
			AvailableMoves = append(AvailableMoves, L3)
		}
		if pc.checkEmptyCell(board, currPos, L4) {
			AvailableMoves = append(AvailableMoves, L4)
		}
		/* if pc.checkEmptyCell(board, currPos, 8) {
			fmt.Printf("OkL5")
			if pc.checkEmptyCell(board, currPos+8, -2) {
				fmt.Printf("OkL5_2")
				AvailableMoves = append(AvailableMoves, L5)
			}
		}
		if pc.checkEmptyCell(board, currPos, L6) {
			AvailableMoves = append(AvailableMoves, L6)
		}
		if pc.checkEmptyCell(board, currPos, 8) {
			if pc.checkEmptyCell(board, currPos+8, 8) {
				if pc.checkEmptyCell(board, currPos+8+8, -1) {
					AvailableMoves = append(AvailableMoves, L7)
				}
			}
		}
		if pc.checkEmptyCell(board, currPos, 8) {
			if pc.checkEmptyCell(board, currPos+8, 8) {
				if pc.checkEmptyCell(board, currPos+8+8, 1) {
					AvailableMoves = append(AvailableMoves, L8)
				}
			}
		} */
		break
	}

}

func (pc *Piece) checkEmptyCell(board []Cell, currPos int, checkPos int) bool {
	// Out of bounds
	if (currPos+checkPos) <= 0 && (currPos+checkPos) >= 63 {
		return false
	}
	// Left border check
	if currPos%8 == 0 && checkPos < 0 && checkPos > -8 {
		return false
	}
	// Right border check
	if currPos%8 == 7 && checkPos > 0 && checkPos < 8 {
		return false
	}
	// Prevent wrapping on right side of board
	if (currPos+checkPos)%8 < currPos%8 && checkPos > 0 && checkPos < 7 {
		//fmt.Printf("oops:%d\n", checkPos)
		return false
	}
	// Prevent wrapping on left side of board
	if (currPos+checkPos)%8 >= currPos%8 && checkPos < 0 && checkPos > -7 {
		return false
	}
	if board[currPos+checkPos].getPieceColor() == NONE {
		return true
	}
	return false
}

// Cell : individual square in chessboard
type Cell struct {
	CellPiece         Piece
	CellOriginalColor Color
	CellCurrentColor  Color
}

func (cl *Cell) setPiece(p ChessPiece, c Color) {
	cl.CellPiece.PieceType = p
	cl.CellPiece.PieceColor = c
}

func (cl *Cell) getPieceColor() Color {
	return cl.CellPiece.PieceColor
}

func (cl *Cell) getCellColor() Color {
	return cl.CellCurrentColor
}

func drawBoard() {
	y := 1
	bg := BG_CYAN
	fg := FG_CYAN
	fmt.Printf("   ")
	// Print X Coords
	for i := 1; i <= 8; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	fmt.Printf("  ")
	for i := 1; i <= 8; i++ {
		fmt.Printf("--")
	}
	fmt.Println()
	fmt.Printf("\0337")
	for i := 0; i < 64; i++ {
		if i == 0 {
			fmt.Printf("%d |", y)
			y++
		}
		// Print Y Coords
		if i != 0 && i%8 == 0 {
			fmt.Println()
			fmt.Printf("\033[0m%d |", y)
			y++
		}
		if (i+y)%2 == 0 {
			bg = BG_CYAN
		} else {
			bg = BG_YELLOW
		}
		if Board[i].CellPiece.PieceColor == BLACK {
			fg = FG_CYAN
		} else {
			fg = FG_YELLOW
		}
		fmt.Printf("\033[%d;%dm%s\033[0m", bg, fg, unicodes[Board[i].CellPiece.PieceType])
	}
	fmt.Println()
}

func initBoardState() {

	unicodes[0] = "  "
	unicodes[1] = "\u265F "
	unicodes[2] = "\u265C "
	unicodes[3] = "\u265E "
	unicodes[4] = "\u265D "
	unicodes[5] = "\u265A "
	unicodes[6] = "\u265B "

	// Ensure empty board
	for i := 0; i < 64; i++ {
		Board[i].setPiece(EMPTY, NONE)
	}
	// Pawn
	for i := 0; i < 8; i++ {
		Board[i+8].setPiece(PAWN, BLACK)
		Board[55-i].setPiece(PAWN, WHITE)
	}
	// Rook
	Board[0].setPiece(ROOK, BLACK)
	Board[7].setPiece(ROOK, BLACK)
	Board[56].setPiece(ROOK, WHITE)
	Board[63].setPiece(ROOK, WHITE)
	// Knight
	Board[1].setPiece(KNIGHT, BLACK)
	Board[6].setPiece(KNIGHT, BLACK)
	Board[57].setPiece(KNIGHT, WHITE)
	Board[62].setPiece(KNIGHT, WHITE)
	// Bishop
	Board[2].setPiece(BISHOP, BLACK)
	Board[5].setPiece(BISHOP, BLACK)
	Board[58].setPiece(BISHOP, WHITE)
	Board[61].setPiece(BISHOP, WHITE)
	// Queen
	Board[3].setPiece(QUEEN, BLACK)
	Board[59].setPiece(QUEEN, WHITE)
	// King
	Board[4].setPiece(KING, BLACK)
	Board[60].setPiece(KING, WHITE)

	Board[37].setPiece(KNIGHT, WHITE)
}

func choosePiece() int {
	fmt.Print("X (1 - 8): ")
	var x int
	var y int
	var pos = -1
	_, err := fmt.Scanf("%d", &x)
	fmt.Print("Y (1 - 8): ")
	_, err = fmt.Scanf("%d", &y)
	if err == nil {
		pos = (8 * (y - 1)) + (x - 1)
	}
	return pos
}

func main() {
	initBoardState()
	drawBoard()
	piece := choosePiece()
	Board[piece].CellPiece.findMoves(Board, piece)
	fmt.Println(AvailableMoves)
}
