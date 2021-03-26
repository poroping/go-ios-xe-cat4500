package main

type L3Vlan struct {
	NedVlan struct {
		Name        int    `json:"name"`
		Description string `json:"description"`
		IP          struct {
			Vrf struct {
			} `json:"vrf"`
			Address struct {
				Primary struct {
					Address string `json:"address"`
					Mask    string `json:"mask"`
				} `json:"primary"`
			} `json:"address"`
		} `json:"ip"`
	} `json:"ned:Vlan"`
}
