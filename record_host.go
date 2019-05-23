package infoblox

import "fmt"

type RecordHostObject struct {
	Object
	Comment         string         `json:"comment,omitempty"`
	ConfigureForDNS bool           `json:"configure_for_dns,omitempty"`
	Ipv4Addrs       []HostIpv4Addr `json:"ipv4addrs,omitempty"`
	//Ipv6Addrs       []HostIpv6Addr `json:"ipv6addrs,omitempty"`
	Name string `json:"name,omitempty"`
	Ttl  int    `json:"ttl,omitempty"`
	View string `json:"view,omitempty"`
}

type HostIpv4Addr struct {
	Object           `json:"-"`
	ConfigureForDHCP bool   `json:"configure_for_dhcp"`
	Host             string `json:"host,omitempty"`
	Ipv4Addr         string `json:"ipv4addr,omitempty"`
	MAC              string `json:"mac,omitempty"`
}

func (c *Client) RecordHost() *Resource {
	return &Resource{
		conn:       c,
		wapiObject: "record:host",
	}
}

func (c *Client) RecordHostObject(ref string) *RecordHostObject {
	host := RecordHostObject{}
	host.Object = Object{
		Ref: ref,
		r:   c.RecordHost(),
	}
	return &host
}

func (c *Client) GetRecordHost(ref string, opts *Options) (*RecordHostObject, error) {
	resp, err := c.RecordHostObject(ref).get(opts)
	if err != nil {
		return nil, fmt.Errorf("Could not get created host record: %s", err)
	}

	var out RecordHostObject
	err = resp.Parse(&out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
