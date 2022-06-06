package colors

func Black(s string) string {
	return "\x1b[30m" + s + "\x1b[0m"
}

func Red(s string) string {
	return "\x1b[31m" + s + "\x1b[0m"
}

func Green(s string) string {
	return "\x1b[32m" + s + "\x1b[0m"
}

func Yellow(s string) string {
	return "\x1b[33m" + s + "\x1b[0m"
}

func Blue(s string) string {
	return "\x1b[34m" + s + "\x1b[0m"
}

func Magenta(s string) string {
	return "\x1b[35m" + s + "\x1b[0m"
}

func Cyan(s string) string {
	return "\x1b[36m" + s + "\x1b[0m"
}

func White(s string) string {
	return "\x1b[37m" + s + "\x1b[0m"
}

func Gray(s string) string {
	return "\x1b[90m" + s + "\x1b[0m"
}

func Grey(s string) string {
	return "\x1b[90m" + s + "\x1b[0m"
}
