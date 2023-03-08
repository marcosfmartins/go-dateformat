package dateformat

import "io"

type reader struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func newReader(input string) *reader {
	r := &reader{input: input}
	r.readChar()

	return r
}

func (r *reader) readChar() {
	if r.readPosition >= len(r.input) {
		r.ch = 0
	} else {
		r.ch = r.input[r.readPosition]
	}

	r.position = r.readPosition
	r.readPosition += 1
}

func (r *reader) peekChar() byte {
	if r.readPosition >= len(r.input) {
		return 0
	} else {
		return r.input[r.readPosition]
	}
}

func isNumber(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func isPrefix(ch byte) bool {
	if ch == '-' || ch == ':' || isNumber(ch) {
		return true
	}
	return false
}

func (r *reader) NextToken() (string, error) {
	if r.ch == 0 {
		return "", io.EOF
	}

	result := []byte{r.ch}

	if r.ch == '%' {
		r.readChar()

		if isLetter(r.ch) {
			result = append(result, r.ch)
		}

		if isPrefix(r.ch) && isLetter(r.peekChar()) {
			result = append(result, r.ch)
			r.readChar()
			result = append(result, r.ch)
		}
	}

	r.readChar()

	return string(result), nil
}
