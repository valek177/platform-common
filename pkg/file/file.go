package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func GetReaderFromStdin() *bufio.Reader {
	return bufio.NewReader(os.Stdin)
}

func GetWriterToStdout() *bufio.Writer {
	return bufio.NewWriter(os.Stdout)
	// need flush defer out.Flush()
}

func GetReaderFromFile(filepath string) *bufio.Reader {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("error while open file %v", err)
	}
	return bufio.NewReader(f)
}

// n is count of lines
func GetLinesFromFile(filepath string, n int) []string {
	in := GetReaderFromFile(filepath)
	if n == 0 {
		res := make([]string, 0)
		for {
			line, err := in.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				} else {
					log.Fatalf("error while reading string %v", err)
				}
			}

			res = append(res, line)
		}
		return res
	}
	res := make([]string, n)
	for i := 0; i < n; i++ {
		line, err := in.ReadString('\n')
		if err != nil {
			log.Fatalf("error while reading string %v", err)
		}

		res[i] = line
		// if need parse ints in line: "0 1 2 3 4" for example
		// nums := strings.Split(line[0:len(line)-1], " ")
	}
	return res
}

func main() {
	// example
	// data := GetLinesFromFile("day2.input", 0)
	// 40 42 45 46 49 47
	// 65 66 68 71 72 72
	// 44 46 49 52 55 59
	// 62 63 66 68 71 74 80
	// 20 23 25 24 26
	// 37 38 35 38 39 38

	// fmt.Println("data is", data[0:10])
	// datais [40 42 45 46 49 47
	//  65 66 68 71 72 72
	//  44 46 49 52 55 59
	//  62 63 66 68 71 74 80 ]
}
