package console

import (
	"fmt"
	"github.com/fatih/color"
	"net"
)

type InfoPrint struct{}

func NewInfoPrint() *InfoPrint {
	return &InfoPrint{}
}

func (i *InfoPrint) Listen(port int) {
	c := color.New(color.FgMagenta)
	hiMagenta := color.New(color.FgHiMagenta)
	fmt.Println(c.Sprint("   ______              "))
	fmt.Println(c.Sprint("  / __/ /__ ___ _  ___ "))
	fmt.Println(c.Sprint(" / _// / _ `/  ' \\/ -_)"))
	fmt.Println(c.Sprint("/_/ /_/\\_,_/_/_/_/\\__/ "))
	fmt.Println(hiMagenta.Sprintf("Server listening on port %v", port))
	fmt.Println(hiMagenta.Sprintf("\t↳ Local: http://127.0.0.1:%v", port))
	netIp, err := i.localIP()
	if err == nil {
		fmt.Println(hiMagenta.Sprintf("\t↳ Network: http://%v:%v", netIp, port))
	}
	fmt.Println()
}

func (*InfoPrint) localIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP, nil
}
