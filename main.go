package main

import (
	"log"
	"os"
)

var (
	codeMap map[byte]int32 = map[byte]int32{
		40: 1,
		41: -1,
	}
)

func main() {
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
	var position int32 = 0
	reported := false
	for index, item := range buffer {
		position += codeMap[item]
		// Part 2
		if reported != true && position == -1 {
			log.Println("Part 2 answer:", index+1)
			reported = true
		}
	}
	// Part 1
	log.Println("Part 1 answer:", position)
}

func init() {
	log.SetOutput(os.Stdout)
}
