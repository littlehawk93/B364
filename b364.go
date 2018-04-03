package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

type parseFunc func(*bufio.Reader) (string, int, int, error)

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

func parse(fn parseFunc, reader *bufio.Reader, err error) (string, int, int, error) {

	if err != nil {
		return 0, 0, err
	}

	return fn(reader)
}

func parseExact(token string, reader *bufio.Reader, err error) (string, int, int, error) {

	if err != nil {
		return "", 0, 0, err
	}

	chars := 9

	for _, strRune := range token {
		r, _, err := reader.ReadRune()

		if err != nil {
			return "", 0, chars, err
		}

		chars++

		if r != strRune {
			return "", 0, chars, fmt.Errorf("Invalid Token. Expecting '%s'", token)
		}
	}

	return token, 0, chars, nil
}

func program(reader *bufio.Reader) (string, int, int, error) {

	lines := 0
	chars := 0

	return lines, chars, nil
}

func function(reader *bufio.Reader) (string, int, int, error) {
	
}

func block(reader *bufio.Reader) (string, int, int, error) {
	
}

func statement(reader *bufio.Reader) (string, int, int, error) {
	
}

func instruction(reader *bufio.Reader) (string, int, int, error) {
	
}

func while(reader *bufio.Reader) (string, int, int, error) {
	
}

func for(reader *bufio.Reader) (string, int, int, error) {
	
}

func if(reader *bufio.Reader) (string, int, int, error) {
	
}

func else(reader *bufio.Reader) (string, int, int, error) {
	
}

func assignment(reader *bufio.Reader) (string, int, int, error) {

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

func modification(reader *bufio.Reader) (string, int, int, error) {
	
}

func print(reader *bufio.Reader) (string, int, int, error) {
	
}

func call(reader *bufio.Reader) (string, int, int, error) {
	
}

func comment(reader *bufio.Reader) (string, int, int, error) {

	r, _, err := reader.ReadRune()

	if err != nil {
		return 0, 0, err
	}

	if r == '#' {

		chars := 1

		_, commentChars, err := parse(commentText, reader, err)

		chars += commentChars

		if err != nil {
			return 0, chars, err
		}
	}
	else
	{
		reader.UnreadRune()
		return 0, 0, errors.New("Expected '#' to begin comment")
	}

	return 0, chars, nil
}

func halt(reader *bufio.Reader) (string, int, int, error) {

}

func expression(reader *bufio.Reader) (string, int, int, error) {

	chars := 0

	return 0, chars, nil
}

func xorTerm(reader *bufio.Reader) (string, int, int, error) {

	chars := 0

	return 0, chars, nil
}

func andTerm(reader *bufio.Reader) (string, int, int, error) {

	chars := 0

	return 0, chars, nil
}

func addTerm(reader *bufio.Reader) (string, int, int, error) {

	chars := 0

	return 0, chars, nil
}

func multiplyTerm(reader *bufio.Reader) (string, int, int, error) {

	chars := 0

	return 0, chars, nil
}

func baseTerm(reader *bufio.Reader) (string, int, int, error) {

	chars := 0

	return 0, chars, nil
}

func literal(reader *bufio.Reader) (string, int, int, error) {

}

func commentText(reader *bufio.Reader) (string, int, int, error) {

	chars := 0

	for true {

		r, _, err = reader.ReadRune()

		if err != nil {
			return 0, chars, err
		}

		if r == '\n' {
			err = reader.UnreadRune()
			return 1, chars, nil
		}

		chars++
	}

	return 0, chars, err
}

func identifier(reader *bufio.Reader) (string, int, int, error) {

	chars := 0

	var buf bytes.Buffer

	result, _, tmpChars, err := parse(letter, reader, nil)

	if err != nil {
		return "", 0, tmpChars, errors.New("Identifier")
	}

	buf.WriteString(result)
	chars += tmpChars

	for true {

		result, _, tmpChars, err = parse(alphaNumeric, reader, nil)

		if err != nil {
			break
		}

		buf.WriteString(result)
		chars += tmpChars
	}

	err = reader.UnreadRune()
	return buf.String() 0, chars, err
}

func number(reader *bufio.Reader) (string, int, int, error) {

	chars := 0
	var buf bytes.Buffer

	r, _, err := reader.ReadRune()

	if err != nil {
		return "", 0, 0, err
	}

	if r != '-' {
		err = reader.UnreadRune()

		if err != nil {
			return "", 0, 0, err
		}
	}

	result, _, tmpChars, err := parse(nonZeroDigit, reader, nil)

	if err != nil {
		return result, 0, tmpChars, errors.New("Number")
	}

	buf.WriteString(result)
	chars += tmpChars

	for true {

		result, _, tmpChars, err := parse(digit, reader, nil)

		if err != nil {
			break
		}

		buf.WriteString(result)
		chars += tmpChars
	}

	err = reader.UnreadRune()
	return buf.String(), 0, chars, err
}

func alphaNumeric(reader *bufio.Reader) (string, int, int, error) {

	result, _, _, err := parse(digit, reader, nil)

	if err != nil {

		result, _, _, err = parse(letter, reader, nil)

		if err != nil {
			return "", 0, 0, errors.New("Alpha Numeric")
		}
	}

	return result, 0, 1, nil
}

func boolean(reader *bufio.Reader) (string, int, int, error) {

	r, _, err := reader.ReadRune()

	if err != nil {
		return "", 0, 0, err
	}

	err = reader.UnreadRune()

	if err != nil {
		return "", 0, 0, err
	}

	if r == 'T' {
		return parseExact("TRUE", reader, err)
	} else if r == 'F' {
		return parseExact("FALSE", reader, err)
	}

	return "", 0, 0, errors.New("Boolean Literal")
}

func digit(reader *bufio.Reader) (string, int, int, error) {

	result, _, _, err := parse(nonZeroDigit, reader, nil)

	if err != nil {

		r, _, err := reader.ReadRune()

		if err != nil {
			return "", 0, 0, err
		}

		if r == '0' {
			return "0", 0, 1, nil
		} else {
			return "", 0, 0, errors.New("Digit")
		}
	}

	return result, 0, 1, err
}

func nonZeroDigit(reader *bufio.Reader) (string, int, int, error) {
	
	r, _, err := reader.ReadRune()

	if !unicode.IsDigit(r) || r == '0' {
		reader.UnreadRune()
		return 0, 0, errors.New("Non-Zero Digit")
	}

	return 0, 1, nil
}

func letter(reader *bufio.Reader) (string, int, int, error) {

	r, _, err := reader.ReadRune()

	if !unicode.IsUpper(r) {
		reader.UnreadRune()
		return 0, 0, errors.New("Letter")
	}

	return 0, 1, nil
}

func allWhitespace(reader *bufio.Reader) (string, int, int, error) {

	lines := 0
	chars := 0

	for true {

		r, _, err := reader.ReadRune()

		if err != nil {
			return lines, chars, err
		}

		if r == '\n' {
			chars++
			lines++
		} else {
			err = r.UnreadRune()

			if err != nil {
				return lines, chars, err
			}

			_, whitespaceChars, err := parse(whitespace, reader, err)

			chars += whitespaceChars

			if err != nil {
				return lines, chars, err
			}

			if whiteSpaceChars == 0 {
				break
			}
		}
	}

	return lines, chars, nil
}

func whitespace(reader *bufio.Reader) (string, int, int, error) {

	var buf bytes.Buffer

	chars := 0

	for true {

		r, _, err := reader.ReadRune()

		if err != nil {
			return buf.String(), 0, chars, err
		}

		if unicode.IsSpace(r) && r != '\n' {
			chars++
			buf.WriteRune(r)
		} else {
			reader.UnreadRune()
			break
		}
	}

	return buf.String(), 0, chars, nil
}
