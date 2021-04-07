package models

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
