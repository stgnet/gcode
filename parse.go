package gcode

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
)

// ParseFile will parse a whole G-Code file from the passed reader.
func ParseFile(r io.Reader) (*File, error) {
	s := bufio.NewScanner(r)

	// prepare file
	file := &File{}

	// read line by line
	for s.Scan() {
		// parse lin
		line, err := ParseLine(s.Text())
		if err != nil {
			return file, err
		}

		// add line
		file.Lines = append(file.Lines, line)
	}

	// check error
	if err := s.Err(); err != nil {
		return file, err
	}

	return file, nil
}

// ParseLine will parse the specified string as a line of G-Codes.
func ParseLine(s string) (Line, error) {
	// prepare line
	l := Line{Codes: GCodes{}}

	// extract line comment
	if i := strings.Index(s, ";"); i >= 0 {
		// save comment
		l.Comment = s[i+1:]

		// reset string
		s = strings.TrimSpace(s[:i])
	}

	// check string
	if s == "" {
		return l, nil
	}

	// parse line
	for s != "" {
		// check for word comment
		if strings.HasPrefix(s, "(") {
			if i := strings.Index(s, ")"); i >= 0 {
				// save comment
				// c.Comment = s[1:i]

				// reset string
				s = strings.TrimSpace(s[i+1:])

				// add code
				// l.Codes = append(l.Codes, c)
				l.Comment = l.Comment + s[0:i+1]

				// go on
				continue
			} else {
				return l, errors.New("missing ) for word comment")
			}
		}

		// check letter
		if !unicode.IsUpper(rune(s[0])) {
			return l, errors.New("expected uppercase letter to begin word")
		}

		// get word and reset string
		var w string
		if i := strings.Index(s, " "); i >= 0 {
			w = s[:i]
			s = strings.TrimSpace(s[i+1:])
		} else {
			w = s
			s = ""
		}

		// check length
		if len(w) < 2 {
			return l, errors.New("expected a word to have at least a length of two")
		}

		// extract letter
		letter := w[0]
		w = w[1:]

		// parse value
		f, err := strconv.ParseFloat(w, 64)
		if err != nil {
			return l, err
		}
		_, exists := l.Codes[byte(letter)]
		if exists {
			return l, errors.New("letter occurs twice on line")
		}

		if letter < 'A' && letter > 'Z' {
			return l, errors.New("unexpected letter not A-Z")
		}

		// add code
		fmt.Printf("l = %#v\n", l)
		l.Codes[byte(letter)] = f
	}

	return l, nil
}
