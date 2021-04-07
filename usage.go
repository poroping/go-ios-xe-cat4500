package main

import (
	"fmt"
	"go-ios-xe-cat4500/client"
)

func main() {
	var HostURL string = "http://172.21.21.29"
	var username string = "cisco"
	var password string = "cisco"
	// intf := L2VlanNoId{NedVlanListNoId: NedVlanListNoId{
	// 	Name: "vd-TEST",
	// 	// ID:   2301,
	// },
	// }
	// intf := L2Vlan{NedVlanList: NedVlanList{
	// 	// ID:   2302,
	// 	Name: "REST-DEVYx",
	// },
	// }
	c, err := client.NewClient(HostURL, username, password)
	// resp, err2 := c.ReadL3Vlan(2301)
	// resp, err2 := c.UpdateL3Vlan(2302, intf)
	// resp, err2 := c.CreateL3Vlan(2302, intf)
	// resp, err2 := c.DeleteL3Vlan(2302)
	// resp, err2 := c.ListL3Vlan()
	// resp, err2 := c.ListBgpNeighbor()
	// resp, err2 := c.ReadBgpNeighbor("10.25.0.1")
	// resp, err2 := c.ReadBgpNeighborConfig("10.25.0.1")
	// resp, err2 := c.ListBgpNeighborConfig()
	// resp, err2 := c.ListBgpNeighborPrefixList("10.25.0.1")
	resp, err2 := c.ReadBgpNeighborPrefixList("10.25.0.1", 43892, "in")

	if err != nil {
		fmt.Println("error", err)
	}
	if err2 != nil {
		fmt.Println("err2", err2)
	}

	fmt.Println(resp)
}
