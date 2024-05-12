package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
)

//	func loadQuestion() []byte {
//		file, err := os.Open("code.txt")
//		if err != nil {
//			log.Panicln("open file with error", err)
//		}
//		stat, err := file.Stat()
//		if err != nil {
//			log.Panicln("open file with error", err)
//		}
//		buffer := make([]byte, stat.Size())
//		file.Read(buffer)
//		file.Close()
//		return buffer
//	}
func getHexString(bytesl []byte) string {

	newkey := new(big.Int).SetBytes(bytesl)

	base16str := fmt.Sprintf("%X", newkey)

	return base16str
}
func main() {

	salt := []byte("ckczppom")
	// correctAnswer := make([]byte, 5, 32)
	found5 := false
	found6 := false
	for i := 0; true; i++ {
		code := append(salt, []byte(strconv.Itoa(i))...)
		hash := md5.Sum(code)
		nonZeroSize := len(getHexString(hash[:]))
		if nonZeroSize == 27 && !found5 {
			log.Println("Part 1 answer is", i)
			found5 = true
		} else if nonZeroSize == 26 && !found6 {
			log.Println("Part 2 answer is", i)
			found6 = true
		}
		if found5 && found6 {
			break
		}
	}
}

func init() {
	log.SetOutput(os.Stdout)
}
