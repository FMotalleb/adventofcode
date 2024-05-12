package main

import (
	"log"
	"os"
)

var (
	operators map[byte][2]int = map[byte][2]int{
		94:  {1, 0},
		62:  {0, 1},
		118: {-1, 0},
		60:  {0, -1},
	}
)

func loadQuestion() []byte {
	file, err := os.Open("code.txt")
	if err != nil {
		log.Panicln("open file with error", err)
	}
	stat, err := file.Stat()
	if err != nil {
		log.Panicln("open file with error", err)
	}
	buffer := make([]byte, stat.Size())
	file.Read(buffer)
	file.Close()
	return buffer
}

func main() {
	resultTable := map[[2]int]int{
		{0, 0}: 1,
	}
	santaPos := [2]int{0, 0}

	//Part 2
	santa2Pos := [2]int{0, 0}
	moveSanta := true

	// To repeat part 1 set isPart1 to true
	isPart1 := false
	com := loadQuestion()

	for _, operation := range com {
		operator := operators[operation]
		var reference *[2]int

		if moveSanta {
			reference = &santaPos
		} else {
			reference = &santa2Pos
		}
		moveSanta = !moveSanta || isPart1

		reference[0] -= operator[0]
		reference[1] -= operator[1]

		resultTable[*reference]++
	}
	log.Println("result:", len(resultTable))
}

func init() {
	log.SetOutput(os.Stdout)
}
