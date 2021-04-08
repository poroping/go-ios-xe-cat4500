package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/poroping/go-ios-xe-cat4500/models"
)

func (c *Client) CreateL3Vlan(id int, description interface{}, ip string, mask string) error {

	m := models.L3Vlan{Vlan: models.Vlan{
		Name: id,
	}}

	if description != nil {
		m.Vlan.Description = fmt.Sprintf("%v", description)
	}

	m.Vlan.IP.Address.Primary.Address = ip
	m.Vlan.IP.Address.Primary.Mask = mask

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/restconf/api/running/native/interface/Vlan/%d", c.HostURL, id), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}
	_, err = c.doRequest(req, 201)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ReadL3Vlan(id int) (*models.L3Vlan, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/api/running/native/interface/Vlan/%d", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	res := models.L3Vlan{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateL3Vlan(id int, description interface{}, ip interface{}, mask interface{}) error {

	m := models.L3Vlan{Vlan: models.Vlan{}}

	if description != nil {
		m.Vlan.Description = fmt.Sprintf("%v", description)
	}
	if ip != nil {
		m.Vlan.IP.Address.Primary.Address = fmt.Sprintf("%v", ip)
	}
	if mask != nil {
		m.Vlan.IP.Address.Primary.Mask = fmt.Sprintf("%v", mask)
	}

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/restconf/api/running/native/interface/Vlan/%d", c.HostURL, id), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}
	_, err = c.doRequest(req, 204)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteL3Vlan(id int) ([]byte, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/restconf/api/running/native/interface/Vlan/%d", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}
	_, err = c.doRequest(req, 204)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Client) ListL3Vlan() (*models.L3VlanList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/api/running/native/interface/Vlan/", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	res := models.L3VlanList{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
