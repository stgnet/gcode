package gcode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testGCode = `; Line Comment
G1 ; After Line Comment
G2 M1
G4
G5 X0 Y0
G6 Z12.7
G7 X-0.4 Y0.8
S3000
X56.666
`

func TestParse(t *testing.T) {
	r := strings.NewReader(testGCode)

	file, err := ParseFile(r)
	assert.NoError(t, err)
	assert.Equal(t, &File{
		Lines: []Line{
			{
				Comment: " Line Comment",
				Codes:   GCodes{},
			},
			{
				Comment: " After Line Comment",
				Codes: GCodes{
					byte('G'): 1,
				},
			},
			/*
				{
					Codes: []GCode{
						// {Comment: "Word Comment"},
					},
				},
			*/
			{
				Codes: GCodes{
					byte('G'): 2,
					byte('M'): 1,
				},
			},
			{
				Codes: GCodes{
					byte('G'): 4,
				},
			},
			{
				Codes: GCodes{
					byte('G'): 5,
					byte('X'): 0,
					byte('Y'): 0,
				},
			},
			{
				Codes: GCodes{
					byte('G'): 6,
					byte('Z'): 12.7,
				},
			},
			{
				Codes: GCodes{
					byte('G'): 7,
					byte('X'): -0.4,
					byte('Y'): 0.8,
				},
			},
			{
				Codes: GCodes{
					byte('S'): 3000,
				},
			},
			{
				Codes: GCodes{
					byte('X'): 56.666,
				},
			},
		},
	}, file)
}

func TestParseInvalid(t *testing.T) {
	gCodes := []string{
		"(Invalid Comment", // <- missing end brace
		"g1",               // <- not upper case
		"G",                // <- missing value
		"GF",               // <- invalid value
	}

	for _, gc := range gCodes {
		_, err := ParseFile(strings.NewReader(gc))
		assert.Error(t, err)
	}
}
