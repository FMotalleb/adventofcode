package main

import (
	"log"
	"os"
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

}

func init() {
	log.SetOutput(os.Stdout)
}
