package main

import (
	"math/rand"
	"time"

	tm "github.com/buger/goterm"
)

const x = 140
const y = 350

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
var outputMap = map[byte]string{
	0: " ",
	1: "▉",
}

func main() {

	old := InitializeBoard(x, y)

	//Single Glider
	//old[0][1], old[1][2], old[2][0], old[2][1], old[2][2] = 1, 1, 1, 1, 1

	tm.Clear()

	for {
		tm.MoveCursor(1, 1)
		for i := range old {
			tm.Printf("%d: %s", (i + 1), RenderIt(old[i]))
			tm.Flush()

		}
		time.Sleep(500 * time.Millisecond)
		old = ComputeNewState(old)

	}
}

//RenderIt (input [][]byte)
func RenderIt(a []byte) string {
	var out string
	for i := range a {
		out += outputMap[a[i]]
	}
	return string(out)
}

//ComputeNewState takes the old state and computes the new
func ComputeNewState(a [][]byte) [][]byte {
	new := InitializeZero(x, y)
	for i := range a {
		for j := range a[i] {
			state := a[i][j]
			sum := GetSum(a, i, j)
			if state == 0 && sum == 3 {
				new[i][j] = 1
			} else if state == 1 && (sum != 2 && sum != 3) {
				new[i][j] = 0
			} else {
				new[i][j] = a[i][j]
			}
		}
	}
	return new
}

func GetSum(a [][]byte, row int, col int) int {
	sum := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			sum += int(a[(row+i+x)%x][(col+j+y)%y])
		}
	}
	sum -= int(a[row][col])
	return sum
}

//InitializeBoard initializes a GoL board of dimensions x and y filled with random
func InitializeBoard(x int, y int) [][]byte {
	output := make([][]byte, x)
	for i := range output {
		output[i] = make([]byte, y)
	}
	for i := range output {
		for j := range output[i] {
			output[i][j] = byte(seededRand.Intn(2))
		}
	}
	return output
}

//InitializeBoard initializes a GoL board of dimensions x and y filled with zeros
func InitializeZero(x int, y int) [][]byte {
	output := make([][]byte, x)
	for i := range output {
		output[i] = make([]byte, y)
	}
	for i := range output {
		for j := range output[i] {
			output[i][j] = 0
		}
	}
	return output
}
