package on99

import "path/filepath"

func SimplifyPath1(path string) string {
	return filepath.Clean(path)
}
