package main

import "fmt"
import "math/rand"

var FREE = 0
var WALL = 1

func main() {

	maze := make([][]int, 10)

	mazeBuild(maze, 0, 9)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print(maze[i][j])
		}
		fmt.Println()
	}

}

func mazeBuild(maze [][]int, topLeftPos int, bottomRightPos int) {
	subMazeWidth := bottomRightPos - topLeftPos + 1
	subMazeHeight := bottomRightPos - topLeftPos + 1
	fmt.Printf("Coordinates of submaze:\n"+
				"Top left: (%d, %d)\n"+
				"Top right: (%d, %d)\n"+
				"Bottom left: (%d, %d)\n"+
				"Bottom right: (%d, %d)\n\n",
				topLeftPos, topLeftPos,
				(topLeftPos + subMazeWidth - 1), topLeftPos,
				(bottomRightPos - subMazeWidth + 1), subMazeHeight-1,
				bottomRightPos, bottomRightPos)
	
	for i := 0; i < subMazeHeight; i++ {
		for j := 0; j < subMazeWidth; j++ {
			maze[i][j] = rand.Int() % 2
		}
	}

	if subMazeHeight == 1 {
		return
	}

	// North West
	mazeBuild(maze, subMazeHeight - (subMazeHeight/2), subMazeWidth - (subMazeWidth/2))
	mazeBuild(maze, subMazeHeight - (subMazeHeight/2), subMazeWidth - (subMazeWidth/2))
	/* row := 0
	column := 0 */
}