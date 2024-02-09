package dev

import (
	"fmt"
	"github.com/fatih/color"
	"golang.org/x/term"
	"math"
	"strings"
	"time"
)

type Request struct{}

func (*Request) Log(m string, p string) {
	var method string

	switch m {
	case "GET":
		method = color.New(color.FgGreen).Sprintf("%v", m)
		break
	case "POST":
		method = color.New(color.FgBlue).Sprintf("%v", m)
		break
	case "PUT":
		method = color.New(color.FgCyan).Sprintf("%v", m)
		break
	case "PATCH":
		method = color.New(color.FgHiYellow).Sprintf("%v", m)
		break
	case "DELETE":
		method = color.New(color.FgRed).Sprintf("%v", m)
		break
	default:
		method = m
		break
	}

	w, _, err := term.GetSize(0)
	t := time.Now().UTC().Format(time.TimeOnly)
	l := w - len(fmt.Sprintf("[ %v ]%v %v", t, m, p)) - 1
	tColor := color.New(color.FgHiBlack).Sprint(t)
	if err != nil || l <= 0 {
		fmt.Printf("[%v] %v %v %v\n", method, tColor, color.New(color.FgHiBlack).Sprint(". . ."), p)
	} else {
		dot := ". "
		dots := strings.Repeat(dot, int(math.Round(float64(l/2))))
		fmt.Printf("[ %v ] %v %v%v\n", method, tColor, color.New(color.FgHiBlack).Sprint(dots), p)

	}
}

func (*Request) FileLog(file string) {
	file = "<public_dir>" + file
	w, _, err := term.GetSize(0)
	l := w - len(fmt.Sprintf("          [ FILE ] %v", file)) - 1
	if err != nil || l <= 0 {
		fmt.Printf("\t↳ [ %v ] %v %v\n", color.New(color.FgHiYellow).Sprint("SERVING FILE"), color.New(color.FgHiBlack).Sprint(". . ."), file)
	} else {
		dot := ". "
		dots := strings.Repeat(dot, int(math.Round(float64(l/2))))
		fmt.Printf("\t↳ [ %v ] %v %v\n", color.New(color.FgHiYellow).Sprint("SERVING FILE"), color.New(color.FgHiBlack).Sprint(dots), file)
	}
}
