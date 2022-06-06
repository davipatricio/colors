package styles

func Dim(s string) string {
	return "\x1b[2m" + s + "\x1b[0m"
}

func Bold(s string) string {
	return "\x1b[1m" + s + "\x1b[0m"
}

func Hidden(s string) string {
	return "\x1b[8m" + s + "\x1b[0m"
}

func Italic(s string) string {
	return "\x1b[3m" + s + "\x1b[0m"
}

func Underline(s string) string {
	return "\x1b[4m" + s + "\x1b[0m"
}

func Inverse(s string) string {
	return "\x1b[7m" + s + "\x1b[0m"
}

func Strikethrough(s string) string {
	return "\x1b[9m" + s + "\x1b[0m"
}

func Reset(s string) string {
	return "\x1b[0m" + s
}
