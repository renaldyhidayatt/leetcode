package grep

import (
	"fmt"
	"os"
	"strings"
)

func Search(pattern string, flags, files []string) []string {
	var fn, fl, fi, fv, fx bool

	for _, f := range flags {
		switch f {
		case "-n":
			fn = true
		case "-l":
			fl = true
		case "-i":
			fi = true
		case "-v":
			fv = true
		case "-x":
			fx = true
		}
	}

	fm := len(files) > 1

	result := []string{}

	for _, filename := range files {
		data, err := os.ReadFile(filename)

		if err == nil {
			lines := strings.Split(string(data), "\n")

			for i, line := range lines {
				tline := line
				tpattern := pattern

				prefix := ""

				if fi {
					tline = strings.ToLower(tline)
					tpattern = strings.ToLower(tpattern)
				}

				if fm {
					prefix += filename + ":"
				}

				if fn {
					prefix += fmt.Sprintf("%v:", i+1)
				}

				found := strings.Contains(tline, tpattern)

				if fx {
					found = tline == tpattern
				}

				if fl && found {
					result = append(result, filename)
					break
				} else if fv && strings.Trim(tline, " ") != "" && !found {
					result = append(result, prefix+line)
				} else if !fv && found {
					result = append(result, prefix+line)
				}
			}
		}
	}

	return result
}
