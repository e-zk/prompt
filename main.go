package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	home    string
	cwd     string
	nocolor bool
)

// Stolen from: https://github.com/icyphox/dotfiles
// Truncates the current working directory:
//   /home/icy/foo/bar -> ~/f/bar
func pwd() string {
	var path string
	if strings.HasPrefix(cwd, home) {
		path = "~" + strings.TrimPrefix(cwd, home)
	} else {
		// If path doesn't contain $HOME, return the
		// entire path as is.
		path = cwd
		return mkcolor(green, path)
	}
	items := strings.Split(path, "/")
	truncItems := []string{}
	for i, item := range items {
		if i == (len(items) - 1) {
			truncItems = append(truncItems, item)
			break
		}
		truncItems = append(truncItems, item[:1])
	}
	return mkcolor(green, filepath.Join(truncItems...))
}

func main() {
	nocolor = os.Getenv("NO_COLOR") == "1"
	home = os.Getenv("HOME")
	cwd, _ = os.Getwd()

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-s":
			fmt.Printf("export PS1=\"\\$(%s) \\$ \";\necho prompt setup;\n", os.Args[0])
			return
		default:
			break
		}
	}
	fmt.Printf("(%s)", pwd())
}
