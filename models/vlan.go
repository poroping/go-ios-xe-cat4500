package models

type Vlan struct {
	VlanList VlanList `json:"ned:vlan-list,omitempty"`
}

type VlanList struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
