package bright

func Black(s string) string {
	return "\x1b[30;1m" + s + "\x1b[0m"
}

func Red(s string) string {
	return "\x1b[31;1m" + s + "\x1b[0m"
}

func Green(s string) string {
	return "\x1b[32;1m" + s + "\x1b[0m"
}

func Yellow(s string) string {
	return "\x1b[33;1m" + s + "\x1b[0m"
}

func Blue(s string) string {
	return "\x1b[34;1m" + s + "\x1b[0m"
}

func Magenta(s string) string {
	return "\x1b[35;1m" + s + "\x1b[0m"
}

func Cyan(s string) string {
	return "\x1b[36;1m" + s + "\x1b[0m"
}

func White(s string) string {
	return "\x1b[37;1m" + s + "\x1b[0m"
}
