package go-ios-xe-cat4500

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

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
