package hard

import "strings"

/*
You own a Goal Parser that can interpret a string command.
The command consists of an alphabet of "G", "()" and/or "(al)" in some order.
The Goal Parser will interpret
"G" as the string "G", "()" as the string "o", and "(al)" as the string "al".
The interpreted strings are then concatenated in the original order.

For example:
input: "G()(al)
output: Goal

input: G()()()()(al)
output: Gooooal

input: (al)G(al)()()G
output: alGalooG
*/
func GoalParsers(strReader *strings.Reader) string {

	var (
		readerBuf = make([]byte, 1024)
	)

	_, err := strReader.Read(readerBuf)
	if err != nil {
		panic(err)
	}
	parsedString := parse(string(readerBuf))
	return parsedString
}

func parse(word string) string {
	output := ""
	start := -1

	for i, char := range word {
		if char == 'G' {
			if start != -1 {
				return ""
			}
			output += "G"
		} else if char == '(' {
			if start != -1 {
				return ""
			}
			start = i
		} else if char == ')' {
			if start == -1 {
				return ""
			}
			token := word[start : i+1]
			if token == "()" {
				output += "o"
			} else if token == "(al)" {
				output += "al"
			} else {
				return ""
			}
			start = -1
		} else if char == 'a' {
			if start == -1 {
				return ""
			}
		} else if char == 'l' {
			if start == -1 || len(output) == 0 || output[len(output)-1] != 'a' {
				return ""
			}
		} else {
			return ""
		}
	}

	if start != -1 {
		return ""
	}

	return output
}
