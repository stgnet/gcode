package gcode

import (
	"fmt"
	"io"
	"strings"
)

// A GCode is either a single G-Code like "X12.3" or an in line comment in the
// form of "(Comment)".
/*
type GCode struct {
	Letter string
	Value  float64
	// Comment string
}
*/
type GCodes map[byte]float64

// String will return a G-Code formatted string.
/*
func (c *GCode) String() string {
	// write G-Code
	return c.Letter + strings.TrimRight(strings.TrimRight(fmt.Sprintf("%f", c.Value), "0"), ".")
}
*/

// A Line consists of multiple G-Codes and a potential line comment.
type Line struct {
	Codes   GCodes
	Comment string
}

var gcOrder = []byte("GMABCDEFHIJKLNOPQRSTUVWXYZ")

// String will return a G-Code formatted string.
func (l *Line) String() string {
	// prepare string
	s := ""

	// write all codes
	for _, i := range gcOrder {
		v, exists := l.Codes[i]
		if !exists {
			continue
		}
		// add space if any codes have been already added
		if len(s) > 0 {
			s = s + " "
		}

		// add string
		s = s + string(i) + strings.TrimRight(strings.TrimRight(fmt.Sprintf("%f", v), "0"), ".")
	}

	// write comment if existing
	if l.Comment != "" {
		// write space if any codes have been written
		if len(l.Codes) > 0 {
			s = s + " "
		}

		// add comment
		s = s + fmt.Sprintf(";%s", l.Comment)
	}

	// add line feed
	s = s + "\n"

	return s
}

// A File contains multiple lines of G-Codes.
type File struct {
	Lines []Line
}

// WriteFile will write the specified G-Code file to the passed writer.
func WriteFile(w io.Writer, f *File) error {
	// generate lines
	for _, l := range f.Lines {
		_, err := io.WriteString(w, l.String())
		if err != nil {
			return err
		}
	}

	return nil
}
