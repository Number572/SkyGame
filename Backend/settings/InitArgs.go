package settings

import (
	"strings"
)

func InitArgs(args []string) {
	var address struct {
		flag bool
		has bool 
	}
	for _, arg := range args {
		switch arg {
			case "-a", "--address":
				address.flag = true
				continue
		}

		switch {
			case address.flag:
				setAddress(arg)
				address.has = true
				address.flag = false
				continue
		}
	}

	if !address.has {
		panic("Address undefined")
	}
}

func setAddress(addr string) {
	splited := strings.Split(addr, ":")
	if len(splited) != 2 {
		panic("Address is corrected")
	}
	Server.Address.IPv4 = splited[0]
	Server.Address.Port = ":" + splited[1]
}
