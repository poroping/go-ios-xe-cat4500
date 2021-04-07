package client

import (
	"encoding/json"
	"fmt"
	"go-ios-xe-cat4500/models"
	"net/http"
	"strings"
)

func (c *Client) CreateBgpNeighbor(id string, as int, newBgpNeighbor models.BgpNeighbor) ([]byte, error) {
	rb, err := json.Marshal(newBgpNeighbor)
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/%d/neighbor/%s", c.HostURL, as, id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	_, err = c.doRequest(req, 201)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Client) ReadBgpNeighbor(id string, as int) (*models.BgpNeighbor, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/%d/neighbor/%s", c.HostURL, as, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	res := models.BgpNeighbor{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateBgpNeighbor(id string, as int, updateBgpNeighbor models.BgpNeighbor) ([]byte, error) {
	rb, err := json.Marshal(updateBgpNeighbor)
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/%d/neighbor/%s", c.HostURL, as, id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	_, err = c.doRequest(req, 204)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Client) DeleteBgpNeighbor(id string, as int) ([]byte, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/%d/neighbor/%s", c.HostURL, as, id), nil)
	if err != nil {
		return nil, err
	}
	_, err = c.doRequest(req, 204)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Client) ListBgpNeighbor() (*models.BgpNeighborList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/43892/neighbor/", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	res := models.BgpNeighborList{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
