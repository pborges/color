package color

type Color []byte

func (c Color)String() string {
	return string(c)
}

var Default Color = Color("\033[0;39m")
var Green Color = Color("\033[0;32m")
var Black Color = Color("\033[0;30m")
var Red Color = Color("\033[0;31m")
var Yellow Color = Color("\033[0;33m")
var Blue Color = Color("\033[0;34m")
var Magenta Color = Color("\033[0;35m")
var Cyan Color = Color("\033[0;36m")
var White Color = Color("\033[0;97m")
