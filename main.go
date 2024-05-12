package main

import (
	"log"
	"os"
	"strings"
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
	data := string(loadQuestion())
	lines := strings.Split(data, "\n")
	okCountPt1 := 0
	okCountPt2 := 0
	for _, line := range lines {
		if IsNicePt1(line) {
			okCountPt1++
		}
		if IsNicePt2(line) {
			okCountPt2++
		}
	}
	log.Println("Part 1:", okCountPt1, ", Part 2:", okCountPt2)
}

func init() {
	log.SetOutput(os.Stdout)
}
