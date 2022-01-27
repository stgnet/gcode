package gcode

// StripComments will remove all inline and line comments from file.
func StripComments(f *File) {
	cl := 0
	for i := range f.Lines {
		j := i - cl
		l := f.Lines[j]

		/*
			cd := 0
			for ii := range l.Codes {
				jj := ii - cd
				c := l.Codes[jj]

				// remove codes with comments
				if c.Comment != "" {
					l.Codes = append(l.Codes[:jj], l.Codes[jj+1:]...)
					cd++
					continue
				}
			}
		*/

		// remove lines with comments or no codes
		if l.Comment != "" || len(l.Codes) == 0 {
			f.Lines = append(f.Lines[:j], f.Lines[j+1:]...)
			cl++
			continue
		}

		// update line
		f.Lines[j] = l
	}
}

// OffsetXYZ will offset all X, Y and Z G-Code values by the specified values.
func OffsetXYZ(f *File, x, y, z float64) {
	for _, l := range f.Lines {
		for g, _ := range l.Codes {
			if g == byte('X') {
				l.Codes[g] += x
			} else if g == byte('Y') {
				l.Codes[g] += y
			} else if g == byte('Z') {
				l.Codes[g] += z
			}
		}
	}
}
