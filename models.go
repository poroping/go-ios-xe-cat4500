package go-ios-xe-cat4500

type BgpNeighbor struct {
	Neighbor Neighbor `json:"ned:neighbor"`
}

type BgpNeighborList struct {
	Collection struct {
		Neighbor []Neighbor `json:"ned:neighbor"`
	} `json:"collection"`
}

type Neighbor struct {
	ID       string `json:"id,omitempty"`
	RemoteAs int    `json:"remote-as"`
}

type BgpNeighborConfig struct {
	NeighborConfig NeighborConfig `json:"ned:neighbor"`
}

type BgpNeighborConfigList struct {
	Collection struct {
		NeighborConfig []NeighborConfig `json:"ned:neighbor"`
	} `json:"collection"`
}

type NeighborConfig struct {
	ID        string        `json:"id,omitempty"`
	Activate  []interface{} `json:"activate"`
	Advertise struct {
		AdditionalPaths struct {
		} `json:"additional-paths"`
	} `json:"advertise"`
	Announce struct {
	} `json:"announce"`
	Capability struct {
	} `json:"capability"`
	DefaultOriginate struct {
	} `json:"default-originate"`
	PrefixList []struct {
		Inout string `json:"inout"`
	} `json:"prefix-list"`
	PathAttribute struct {
		Discard struct {
		} `json:"discard"`
		TreatAsWithdraw struct {
		} `json:"treat-as-withdraw"`
	} `json:"path-attribute"`
	RemovePrivateAs     []interface{} `json:"remove-private-as"`
	SoftReconfiguration string        `json:"soft-reconfiguration"`
	SlowPeer            struct {
	} `json:"slow-peer"`
	TranslateUpdate struct {
	} `json:"translate-update"`
}

type BgpNeighborPrefixListList struct {
	Collection struct {
		PrefixList []PrefixList `json:"ned:prefix-list"`
	} `json:"collection"`
}

type BgpNeighborPrefixList struct {
	PrefixList PrefixList `json:"ned:prefix-list"`
}

type PrefixList struct {
	Inout          string `json:"inout"`
	PrefixListName string `json:"prefix-list-name"`
}

type L3VlanList struct {
	Collection struct {
		Vlan []Vlan `json:"ned:Vlan"`
	} `json:"collection"`
}

type L3Vlan struct {
	Vlan Vlan `json:"ned:Vlan"`
}

type Vlan struct {
	Name        int    `json:"name,omitempty"`
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
}

type L2Vlan struct {
	VlanList VlanList `json:"ned:vlan-list"`
}

type L2VlanList struct {
	Collection struct {
		VlanList []VlanList `json:"ned:vlan-list"`
	} `json:"collection"`
}

type VlanList struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
}