package main

import (
	"fmt"
)

func main() {
	var HostURL string = "https://172.29.72.81"
	var username string = "cisco"
	var password string = "cisco"
	intf := SingleInterface{IetfInterfacesInterface: Interface{
		Name:        "Vlan404",
		Description: "golang test",
		Type:        "iana-if-type:l3ipvlan",
		Enabled:     true,
		IetfIPIpv4: IetfIPIpv4{
			Address: []Address{{
				IP:      "192.168.8.1",
				Netmask: "255.255.255.248",
			},
			},
		},
		IetfIPIpv6: IetfIPIpv6{},
	},
	}
	c, err := NewClient(HostURL, username, password)
	// resp, err2 := c.GetInterfaces()
	// resp, err2 := c.GetInterface("Vlan400")
	// resp, err2 := c.CreateInterface(intf)
	resp, err2 := c.EditInterface(intf)
	// resp, err2 := c.DeleteInterface("Vlan404")

	if err != nil {
		fmt.Println("error", err)
	}
	if err2 != nil {
		fmt.Println("err2", err2)
	}

	fmt.Println(resp)
}
