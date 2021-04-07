package client

import (
	"encoding/json"
	"fmt"
	"go-ios-xe-cat4500/models"
	"net/http"
	"strings"
)

func (c *Client) CreateBgpNeighborConfig(id string, as int, newBgpNeighborConfig models.BgpNeighborConfig) ([]byte, error) {
	rb, err := json.Marshal(newBgpNeighborConfig)
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/%d/address-family/no-vrf/ipv4/unicast/neighbor/%s", c.HostURL, as, id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	_, err = c.doRequest(req, 201)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Client) ReadBgpNeighborConfig(id string, as int) (*models.BgpNeighborConfig, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/%d/address-family/no-vrf/ipv4/unicast/neighbor/%s", c.HostURL, as, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	res := models.BgpNeighborConfig{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateBgpNeighborConfig(id string, as int, updateBgpNeighborConfig models.BgpNeighborConfig) ([]byte, error) {
	rb, err := json.Marshal(updateBgpNeighborConfig)
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/%d/address-family/no-vrf/ipv4/unicast/neighbor/%s", c.HostURL, as, id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	_, err = c.doRequest(req, 204)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Client) DeleteBgpNeighborConfig(id string, as int) ([]byte, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/%d/address-family/no-vrf/ipv4/unicast/neighbor/%s", c.HostURL, as, id), nil)
	if err != nil {
		return nil, err
	}
	_, err = c.doRequest(req, 204)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Client) ListBgpNeighborConfig(as int) (*models.BgpNeighborConfigList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/%d/address-family/no-vrf/ipv4/unicast/neighbor/", c.HostURL, as), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	res := models.BgpNeighborConfigList{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
