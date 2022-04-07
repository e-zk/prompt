package main

// color escape codes defined here
const (
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	white   = "\033[37m"

	// reset
	rc = "\033[0m"
)

func mkcolor(color, text string) string {
	if nocolor {
		return text
	} else {
		return color + text + rc
	}
}
