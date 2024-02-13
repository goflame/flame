package utils

import (
	"fmt"
	"github.com/fatih/color"
)

type log struct{}

var Log log

func (l *log) Info(msg ...interface{}) {
	c := color.New(color.FgWhite, color.BgBlue)
	fmt.Printf("%v %v\n", c.Sprint(" INFO "), color.New(color.FgHiBlue).Sprint(msg...))
}

func (l *log) Warn(msg ...interface{}) {
	c := color.New(color.FgWhite, color.BgYellow)
	fmt.Printf("%v %v\n", c.Sprint(" WARN "), color.New(color.FgHiYellow).Sprint(msg...))
}

func (l *log) Error(msg ...interface{}) {
	c := color.New(color.FgWhite, color.BgRed)
	fmt.Printf("%v %v\n", c.Sprint(" ERROR "), color.New(color.FgHiRed).Sprint(msg...))
}

func (l *log) Success(msg ...interface{}) {
	c := color.New(color.FgWhite, color.BgGreen)
	fmt.Printf("%v %v\n", c.Sprint(" SUCCESS "), color.New(color.FgHiGreen).Sprint(msg...))
}
