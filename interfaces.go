package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Interfaces struct {
	IetfInterfacesInterface []Interface `json:"ietf-interfaces:interface"`
}

type SingleInterface struct {
	IetfInterfacesInterface Interface `json:"ietf-interfaces:interface"`
}

type IetfIPIpv4 struct {
	Address []Address `json:"address"`
}

type IetfIPIpv6 struct {
}

type Address struct {
	IP      string `json:"ip"`
	Netmask string `json:"netmask"`
}

type Interface struct {
	Name        string     `json:"name"`
	Description string     `json:"description,omitempty"`
	Type        string     `json:"type"`
	Enabled     bool       `json:"enabled"`
	IetfIPIpv4  IetfIPIpv4 `json:"ietf-ip:ipv4,omitempty"`
	IetfIPIpv6  IetfIPIpv6 `json:"ietf-ip:ipv6,omitempty"`
}

func (c *Client) GetInterfaces() (*Interfaces, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/data/ietf-interfaces:interfaces/interface=", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	interfaces := Interfaces{}
	err = json.Unmarshal(body, &interfaces)
	if err != nil {
		return nil, err
	}

	return &interfaces, nil
}

func (c *Client) GetInterface(name string) (*SingleInterface, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/data/ietf-interfaces:interfaces/interface=%s", c.HostURL, name), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	interf := SingleInterface{}
	fmt.Println(interf)
	err = json.Unmarshal(body, &interf)
	if err != nil {
		return nil, err
	}

	return &interf, nil
}

func (c *Client) CreateInterface(newInterface SingleInterface) ([]byte, error) {
	rb, err := json.Marshal(newInterface)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/restconf/data/ietf-interfaces:interfaces", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	_, err = c.doRequest(req, 201)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Client) EditInterface(editInterface SingleInterface) ([]byte, error) {
	rb, err := json.Marshal(editInterface)
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/restconf/data/ietf-interfaces:interfaces/interface=%s", c.HostURL, editInterface.IetfInterfacesInterface.Name), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	_, err = c.doRequest(req, 204)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *Client) DeleteInterface(name string) ([]byte, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/restconf/data/ietf-interfaces:interfaces/interface=%s", c.HostURL, name), nil)
	if err != nil {
		return nil, err
	}
	_, err = c.doRequest(req, 204)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
