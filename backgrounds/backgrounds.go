package backgrounds

func Black(s string) string {
	return "\x1b[40m" + s + "\x1b[0m"
}

func Red(s string) string {
	return "\x1b[41m" + s + "\x1b[0m"
}

func Green(s string) string {
	return "\x1b[42m" + s + "\x1b[0m"
}

func Yellow(s string) string {
	return "\x1b[43m" + s + "\x1b[0m"
}

func Blue(s string) string {
	return "\x1b[44m" + s + "\x1b[0m"
}

func Magenta(s string) string {
	return "\x1b[45m" + s + "\x1b[0m"
}

func Cyan(s string) string {
	return "\x1b[46m" + s + "\x1b[0m"
}

func White(s string) string {
	return "\x1b[47m" + s + "\x1b[0m"
}
