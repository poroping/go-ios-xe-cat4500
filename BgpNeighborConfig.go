package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type BgpNeighborConfig struct {
	NedNeighborConfig NedNeighborConfig `json:"ned:neighbor"`
}

type BgpNeighborConfigList struct {
	Collection struct {
		NedNeighborConfig []NedNeighborConfig `json:"ned:neighbor"`
	} `json:"collection"`
}

type NedNeighborConfig struct {
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

func (c *Client) CreateBgpNeighborConfig(id string, newBgpNeighborConfig BgpNeighborConfig) ([]byte, error) {
	rb, err := json.Marshal(newBgpNeighborConfig)
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/43892/address-family/no-vrf/ipv4/unicast/neighbor/%s", c.HostURL, id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	_, err = c.doRequest(req, 201)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Client) ReadBgpNeighborConfig(id string) (*BgpNeighborConfig, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/43892/address-family/no-vrf/ipv4/unicast/neighbor/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	res := BgpNeighborConfig{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateBgpNeighborConfig(id string, updateBgpNeighborConfig BgpNeighborConfig) ([]byte, error) {
	rb, err := json.Marshal(updateBgpNeighborConfig)
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/43892/address-family/no-vrf/ipv4/unicast/neighbor/%s", c.HostURL, id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	_, err = c.doRequest(req, 204)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Client) DeleteBgpNeighborConfig(id string) ([]byte, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/43892/address-family/no-vrf/ipv4/unicast/neighbor/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}
	_, err = c.doRequest(req, 204)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Client) ListBgpNeighborConfig() (*BgpNeighborConfigList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/43892/address-family/no-vrf/ipv4/unicast/neighbor/", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	res := BgpNeighborConfigList{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
