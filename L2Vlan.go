package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type L2Vlan struct {
	NedVlanList NedVlanList `json:"ned:vlan-list"`
}

type L2VlanList struct {
	Collection struct {
		NedVlanList []NedVlanList `json:"ned:vlan-list"`
	} `json:"collection"`
}

type NedVlanList struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

func (c *Client) CreateL2Vlan(id int, newL2Vlan L2Vlan) ([]byte, error) {
	rb, err := json.Marshal(newL2Vlan)
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/restconf/api/running/native/vlan/vlan-list/%d", c.HostURL, id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	_, err = c.doRequest(req, 201)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Client) ReadL2Vlan(id int) (*L2Vlan, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/api/running/native/vlan/vlan-list/%d", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	res := L2Vlan{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateL2Vlan(id int, updateL2Vlan L2Vlan) ([]byte, error) {
	rb, err := json.Marshal(updateL2Vlan)
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/restconf/api/running/native/vlan/vlan-list/%d", c.HostURL, id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	_, err = c.doRequest(req, 204)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Client) DeleteL2Vlan(id int) ([]byte, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/restconf/api/running/native/vlan/vlan-list/%d", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}
	_, err = c.doRequest(req, 204)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Client) ListL2Vlan() (*L2VlanList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/api/running/native/vlan/vlan-list/", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	res := L2VlanList{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
