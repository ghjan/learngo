package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

func main() {
	fmt.Println(convertToBin(5), // 101 ->101
		convertToBin(13),
		convertToBin(72387885),
		convertToBin(0),
		convertToBin(-5),
	) // 1011->1101

	printFileLoop("abc.txt")
	forever()
}

func convertToBin(n int) string {
	if n == 0 {
		return "0"
	} else if n < 0 {
		return strconv.FormatInt(int64(n), 2)
	}
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFileLoop(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("abc")
	}
}
