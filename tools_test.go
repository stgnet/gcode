package gcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStripComments(t *testing.T) {
	f := &File{
		Lines: []Line{
			{
				Comment: "foo",
			},
			{
				Codes: GCodes{
					// {Comment: "bar"},
				},
			},
			{
				Codes: GCodes{
					byte('A'): 0,
					// {Comment: "baz"},
				},
			},
		},
	}

	StripComments(f)

	assert.Equal(t, &File{
		Lines: []Line{
			{
				Codes: GCodes{
					byte('A'): 0,
				},
			},
		},
	}, f)
}

func TestOffsetXYZ(t *testing.T) {
	f := &File{
		Lines: []Line{
			{
				Codes: GCodes{
					byte('G'): 1,
				},
			},
			{
				Codes: GCodes{
					byte('X'): 2,
				},
			},
			{
				Codes: GCodes{
					byte('X'): 3,
					byte('Y'): 4,
					byte('Z'): 5,
				},
			},
		},
	}

	OffsetXYZ(f, 1, 2, 3)

	assert.Equal(t, &File{
		Lines: []Line{
			{
				Codes: GCodes{
					byte('G'): 1,
				},
			},
			{
				Codes: GCodes{
					byte('X'): 3,
				},
			},
			{
				Codes: GCodes{
					byte('X'): 4,
					byte('Y'): 6,
					byte('Z'): 8,
				},
			},
		},
	}, f)
}
