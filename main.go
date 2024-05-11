package main

import (
	"log"
	"os"
	"sort"
	"strconv"
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
	lines := strings.Split(string(loadQuestion()), "\n")
	totalArea, totalRibbon := 0, 0
	for _, item := range lines {
		g := parseGift(item)

		totalArea += g.calculateArea()
		totalRibbon += g.calculateRibbon()

	}
	log.Println("Part 1:", totalArea, "and Part 2:", totalRibbon)
}

func init() {
	log.SetOutput(os.Stdout)
}

type gift struct {
	dimensions *[3]int
}

func (g gift) calculateArea() int {
	dims := g.dimensions
	return dims[0]*dims[1] + 2*dims[0]*dims[1] + 2*dims[1]*dims[2] + 2*dims[0]*dims[2]
}
func (g gift) calculateRibbon() int {
	dims := g.dimensions
	return dims[0]*2 + dims[1]*2 + dims[0]*dims[1]*dims[2]
}

func parseGift(data string) gift {
	dimensions := strings.Split(data, "x")
	a, _ := strconv.Atoi(dimensions[0])
	b, _ := strconv.Atoi(dimensions[1])
	c, _ := strconv.Atoi(dimensions[2])
	finalDimensions := &[3]int{
		a, b, c,
	}
	sort.Ints(finalDimensions[:])
	return gift{
		finalDimensions,
	}
}
