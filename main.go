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
	currentPos := [2]int{0, 0}
	com := loadQuestion()
	for _, operation := range com {
		operator := operators[operation]
		currentPos[0] += operator[0]
		currentPos[1] += operator[1]
		resultTable[currentPos]++
	}
	log.Println("all houses", len(resultTable))
}

func init() {
	log.SetOutput(os.Stdout)
}
