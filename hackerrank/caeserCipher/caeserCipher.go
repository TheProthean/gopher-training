package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the caesarCipher function below.
func caesarCipher(s string, k int32) string {
	alphabetSize := 26
	modK := int(k) % alphabetSize
	result := make([]byte, len(s))
	for i := range s {
		if s[i] >= 97 && s[i] <= 122 {
			newB := s[i] + byte(modK)
			if newB > 122 {
				newB -= byte(alphabetSize)
			}
			result[i] = newB
		} else if s[i] >= 65 && s[i] <= 90 {
			newB := s[i] + byte(modK)
			if newB > 90 {
				newB -= byte(alphabetSize)
			}
			result[i] = newB
		} else {
			result[i] = s[i]
		}
	}
	return string(result)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	_, err = strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	s := readLine(reader)

	kTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := caesarCipher(s, k)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
