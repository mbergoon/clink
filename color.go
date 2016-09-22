package main

type ColorCode string

const (
	COLOR_B_BLACK  ColorCode = "\x1b[30;1m"
	COLOR_B_RED    ColorCode = "\x1b[31;1m"
	COLOR_B_GREEN  ColorCode = "\x1b[32;1m"
	COLOR_B_YELLOW ColorCode = "\x1b[33;1m"
	COLOR_B_BLUE   ColorCode = "\x1b[34;1m"
	COLOR_B_MAROON ColorCode = "\x1b[35;1m"
	COLOR_B_CYAN   ColorCode = "\x1b[36;1m"
	COLOR_B_WHITE  ColorCode = "\x1b[37;1m"

	COLOR_BLACK  ColorCode = "\x1b[30m"
	COLOR_RER    ColorCode = "\x1b[31m"
	COLOR_GREEN  ColorCode = "\x1b[32m"
	COLOR_YELLOW ColorCode = "\x1b[33m"
	COLOR_BLUE   ColorCode = "\x1b[34m"
	COLOR_MAROON ColorCode = "\x1b[35m"
	COLOR_CYAN   ColorCode = "\x1b[36m"
	COLOR_WHITE  ColorCode = "\x1b[37m"

	COLOR_NONE ColorCode = "\x1b[0m"
)

func (cc ColorCode) String() string {
	return string(cc)
}

func Color(c ColorCode, s string) string {
	return c.String() + s
}

func ColorClear(c ColorCode, s string) string {
	return c.String() + s + COLOR_NONE.String()
}
