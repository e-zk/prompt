package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	home    string
	cwd     string
	nocolor bool
)

// Truncates the current working directory:
//   /home/icy/foo/bar -> ~/f/bar
// Stolen from: https://github.com/icyphox/dotfiles
func pwd() string {
	var path string
	if strings.HasPrefix(cwd, home) {
		path = "~" + strings.TrimPrefix(cwd, home)
	} else {
		// If path doesn't contain $HOME, return the
		// entire path as is.
		path = cwd
		//return path
	}

	dirs := strings.Split(path, "/")
	for i, d := range dirs {
		if i == 0 || i == len(dirs)-1 {
			continue
		}
		dirs[i] = string(d[0])
	}
	path = strings.Join(dirs, "/")

	return mkcolor(green, path)
}

func printPrompt() (out string) {
	return fmt.Sprintf("(%s)", pwd())
}

func main() {
	nocolor = os.Getenv("NO_COLOR") == "1"
	home = os.Getenv("HOME")
	cwd, _ = os.Getwd()

	if len(os.Args) < 2 {
		print(printPrompt())
		return
	}

	switch os.Args[1] {
	case "-s":
		fmt.Printf("export PS1=\"\\$(%s) \\$ \";\necho prompt setup;\n", os.Args[0])
		return
	default:
		print(printPrompt())
		return
	}
}
