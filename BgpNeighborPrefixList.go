package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

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

func (c *Client) CreateBgpNeighborPrefixList(id string, inout string, newBgpNeighborPrefixList BgpNeighborPrefixList) ([]byte, error) {
	rb, err := json.Marshal(newBgpNeighborPrefixList)
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/43892/address-family/no-vrf/ipv4/unicast/neighbor/%s/prefix-list/%s/", c.HostURL, id, inout), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	_, err = c.doRequest(req, 201)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Client) ReadBgpNeighborPrefixList(id string, inout string) (*BgpNeighborPrefixList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/43892/address-family/no-vrf/ipv4/unicast/neighbor/%s/prefix-list/%s/", c.HostURL, id, inout), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	res := BgpNeighborPrefixList{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateBgpNeighborPrefixList(id string, inout string, updateBgpNeighborPrefixList BgpNeighborPrefixList) ([]byte, error) {
	rb, err := json.Marshal(updateBgpNeighborPrefixList)
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/43892/address-family/no-vrf/ipv4/unicast/neighbor/%s/prefix-list/%s/", c.HostURL, id, inout), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	_, err = c.doRequest(req, 204)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Client) DeleteBgpNeighborPrefixList(id string, inout string) ([]byte, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/43892/address-family/no-vrf/ipv4/unicast/neighbor/%s/prefix-list/%s/", c.HostURL, id, inout), nil)
	if err != nil {
		return nil, err
	}
	_, err = c.doRequest(req, 204)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Client) ListBgpNeighborPrefixList(id string) (*BgpNeighborPrefixListList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/43892/address-family/no-vrf/ipv4/unicast/neighbor/%s/prefix-list/", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	res := BgpNeighborPrefixListList{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
