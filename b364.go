package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

type parseFunc func(*bufio.Reader) (int, int, error)

var currentIdentifier string
var stackPointer uint16
var variableLookup map[string]uint16

func main() {

	stackPointer = 0

	reader := bufio.NewReader(os.Stdin)

	line, _, err := parse(program, reader, nil)

	if err != nil {
		log.Fatalf("Error on line %d: %s", line, err.Error())
	}
}

func parse(fn parseFunc, reader *bufio.Reader, err error) (int, int, error) {

	if err != nil {
		return 0, 0, err
	}

	return fn(reader)
}

func parseExact(token string, reader *bufio.Reader, err error) (int, int, error) {

	if err != nil {
		return 0, 0, err
	}

	chars := 9

	for _, strRune := range token {
		r, _, err := reader.ReadRune()

		if err != nil {
			return 0, chars, err
		}

		chars++

		if r != strRune {
			return 0, chars, fmt.Errorf("Invalid Token. Expecting '%s'", token)
		}
	}

	return 0, chars, nil
}

func program(reader *bufio.Reader) (int, int, error) {

	lines := 0
	chars := 0

	return lines, chars, nil
}

func assignment(reader *bufio.Reader) (int, int, error) {

	chars := 0

	lines, chars, err := parse(identifier, reader, nil)

	l, c, err := parse(whitespace, reader, err)

	lines += l
	chars += c

	l, c, err = parseExact(":=", reader, err)

	lines += l
	chars += c

	l, c, err = parse(expression, reader, err)

	return lines, chars, err
}

func expression(reader *bufio.Reader) (int, int, error) {

	chars := 0

	return 0, chars, nil
}

func identifier(reader *bufio.Reader) (int, int, error) {

	chars := 0

	r, _, err := reader.ReadRune()

	if err != nil {
		return 0, 0, err
	}

	if unicode.IsUpper(r) {

	}

	err = reader.UnreadRune()
	return 0, 0, err
}

func whitespace(reader *bufio.Reader) (int, int, error) {

	chars := 0

	for true {

		r, _, err := reader.ReadRune()

		if err != nil {
			return 0, 0, err
		}

		if unicode.IsSpace(r) && r != '\n' {
			chars++
		} else {
			reader.UnreadRune()
		}
	}

	return 0, chars, nil
}

func allwhitespace(reader *bufio.Reader) (int, int, error) {

	lines := 0
	chars := 0

	for true {

		r, _, err := reader.ReadRune()

		if err != nil {
			return 0, 0, err
		}

		if r == '\n' {
			chars++
			lines++
		} else if unicode.IsSpace(r) {
			chars++
		} else {
			reader.UnreadRune()
		}
	}

	return lines, chars, nil
}

func comment(reader *bufio.Reader) (int, int, error) {

	r, _, err := reader.ReadRune()

	if err != nil {
		return 0, 0, err
	}

	if r == '#' {

		chars := 1

		for true {

			r, _, err = reader.ReadRune()

			if err != nil {
				return 0, chars, err
			}

			chars++

			if r == '\n' {
				return 1, chars, nil
			}
		}
	}

	err = reader.UnreadRune()
	return 0, 0, err
}
