package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type BgpNeighbor struct {
	NedNeighbor NedNeighbor `json:"ned:neighbor"`
}

type BgpNeighborList struct {
	Collection struct {
		NedNeighbor []NedNeighbor `json:"ned:neighbor"`
	} `json:"collection"`
}

type NedNeighbor struct {
	ID       string `json:"id,omitempty"`
	RemoteAs int    `json:"remote-as"`
}

func (c *Client) CreateBgpNeighbor(id string, newBgpNeighbor BgpNeighbor) ([]byte, error) {
	rb, err := json.Marshal(newBgpNeighbor)
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/43892/neighbor/%s", c.HostURL, id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	_, err = c.doRequest(req, 201)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Client) ReadBgpNeighbor(id string) (*BgpNeighbor, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/43892/neighbor/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	res := BgpNeighbor{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateBgpNeighbor(id string, updateBgpNeighbor BgpNeighbor) ([]byte, error) {
	rb, err := json.Marshal(updateBgpNeighbor)
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/43892/neighbor/%s", c.HostURL, id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	_, err = c.doRequest(req, 204)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Client) DeleteBgpNeighbor(id string) ([]byte, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/43892/neighbor/%s", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}
	_, err = c.doRequest(req, 204)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Client) ListBgpNeighbor() (*BgpNeighborList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/43892/neighbor/", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	res := BgpNeighborList{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
