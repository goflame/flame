package console

import (
	"fmt"
	"github.com/fatih/color"
)

type InfoPrint struct{}

func (InfoPrint) Listen(port int) {
	c := color.New(color.FgMagenta)
	fmt.Println(c.Sprint("   ______              "))
	fmt.Println(c.Sprint("  / __/ /__ ___ _  ___ "))
	fmt.Println(c.Sprint(" / _// / _ `/  ' \\/ -_)"))
	fmt.Println(c.Sprint("/_/ /_/\\_,_/_/_/_/\\__/ "))
	fmt.Println(color.New(color.FgHiMagenta).Sprintf("Server listening on port %v\n", port))
}
