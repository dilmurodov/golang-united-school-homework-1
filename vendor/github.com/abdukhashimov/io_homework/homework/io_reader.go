package homework

import (
	"io"
	"math"
	"strings"
)

/*
SeekTillHalfOfString -  contains a code snippet in Go that defines a function called
"SeekTillHalfOfString". The function takes a string reader as input,
seeks to the middle of the string, reads
half of the remaining string, and returns it as a string.
*/
func SeekTillHalfOfString(strReader *strings.Reader) string {
	halfSize := float64(strReader.Size()) / 2.0
	readSize := int64(math.Ceil(float64(halfSize)))
	strReader.Seek(int64(halfSize), io.SeekStart)
	readerBuf := make([]byte, readSize)
	strReader.Read(readerBuf)
	return string(readerBuf)
}

/*
ReaderSplit - contains a code snippet written in Go that
defines a function called ReaderSplit.
The function takes a strings.Reader and an integer n as input,
and splits the contents of the reader into chunks of size n.
The function returns a slice of strings containing the chunks
*/
func ReaderSplit(strReader *strings.Reader, n int) []string {
	readerBuf := make([]byte, n)
	resp := []string{}

	for {
		n, err := strReader.Read(readerBuf)
		if err == io.EOF {
			break
		}

		resp = append(resp, string(readerBuf[:n]))
	}

	return resp
}

/*
You own a Goal Parser that can interpret a string command.
The command consists of an alphabet of "G", "()" and/or "(al)" in some order.
The Goal Parser will interpret
"G" as the string "G", "()" as the string "o", and "(al)" as the string "al".
The interpreted strings are then concatenated in the original order.
*/
func GoalParsers(strReader *strings.Reader) string {
	var (
		readerBuf = make([]byte, 1024)
	)

	n, err := strReader.Read(readerBuf)
	if err != nil {
		return ""
	}

	if n == 0 {
		return ""
	}

	return parse(string(readerBuf))
}

func parse(word string) string {
	var (
		parsing string
	)

	if len(word) == 0 {
		return ""
	}

	for i := 0; i < len(word); i++ {
		if word[i] == 'G' {
			parsing = "G"
		} else if word[i] == '(' {
			parsing += "("
		} else if word[i] == ')' {
			parsing += ")"
		} else if word[i] == 'a' {
			parsing += "a"
		} else if (word[i]) == 'l' {
			parsing += "l"
		}

		if parsing == "G" {
			parsing = ""
			return "G" + parse(word[i+1:])
		} else if parsing == "()" {
			parsing = ""
			return "o" + parse(word[i+1:])
		} else if parsing == "(al)" {
			parsing = ""
			return "al" + parse(word[i+1:])
		}
	}

	return ""
}
