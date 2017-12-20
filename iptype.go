package iputility

import (
	"net"
	"strings"
)

type IpId uint8

type Ip struct {
	Type     IpId
	Endpoint string
	loUint   uint64
	hiUint   uint64
}

const (
	IPTYPE_ADDRESS IpId = iota + 1
	IPTYPE_CIDR
	IPTYPE_RANGE
	IPTYPE_FQDN
	IPTYPE_UNDEFINED
)

func GetType(endpoint string) Ip {

	ip := net.ParseIP(endpoint)
	if ip != nil {
		ipUint := toUint64(ip)
		return Ip{Type: IPTYPE_ADDRESS, Endpoint: endpoint, loUint: ipUint, hiUint: ipUint}
	}

	cidrIP, subnet, cidrErr := net.ParseCIDR(endpoint)
	if cidrErr == nil {
		os, bs := subnet.Mask.Size()

		loUint := toUint64(cidrIP)
		hiUint := loUint + (1 << uint64(32-os)) - 1

		if os == bs {
			return Ip{Type: IPTYPE_ADDRESS, Endpoint: cidrIP.String(), loUint: loUint, hiUint: hiUint}
		}

		return Ip{Type: IPTYPE_CIDR, Endpoint: endpoint, loUint: loUint, hiUint: hiUint}
	}

	ipRange := strings.Split(endpoint, "-")
	if len(ipRange) == 2 {
		ipLo := net.ParseIP(ipRange[0])
		ipHi := net.ParseIP(ipRange[1])
		if ipLo != nil && ipHi != nil {

			loUint := toUint64(ipLo)
			hiUint := toUint64(ipHi)
			if loUint > hiUint {
				return Ip{Type: IPTYPE_UNDEFINED, Endpoint: ""}
			} else if loUint == hiUint {
				return Ip{Type: IPTYPE_ADDRESS, Endpoint: ipRange[0], loUint: loUint, hiUint: hiUint}
			}

			return Ip{Type: IPTYPE_RANGE, Endpoint: endpoint, loUint: loUint, hiUint: hiUint}
		}
	}

	if ip == nil && cidrErr != nil && strings.Index(endpoint, ".") > 0 {
		return Ip{Type: IPTYPE_FQDN, Endpoint: endpoint}
	}

	return Ip{Type: IPTYPE_UNDEFINED, Endpoint: ""}
}

func (t *Ip) GetFirst() string {

	switch t.Type {
	case IPTYPE_ADDRESS:
		return t.Endpoint
	case IPTYPE_CIDR:
		return strings.Split(t.Endpoint, "/")[0]
	case IPTYPE_RANGE:
		return strings.Split(t.Endpoint, "-")[0]
	case IPTYPE_FQDN:
		return t.Endpoint
	}
	return t.Endpoint
}

func (t *Ip) In(t1 Ip) bool {

	if !t.isIpType() || !t1.isIpType() {
		return false
	}

	if (t.loUint >= t1.loUint && t.hiUint < t1.hiUint) ||
		(t.loUint > t1.loUint && t.hiUint <= t1.hiUint) {
		return true
	}

	return false
}

func (t *Ip) Equals(t1 Ip) bool {

	if !t.isIpType() || !t1.isIpType() {
		return false
	}

	if t.loUint == t1.loUint && t.hiUint == t1.hiUint {
		return true
	}

	return false
}

func (t *Ip) isIpType() bool {

	if t.Type != IPTYPE_ADDRESS && t.Type != IPTYPE_CIDR && t.Type != IPTYPE_RANGE {
		return false
	}

	return true
}

func toUint64(ip net.IP) uint64 {

	if len(ip) != net.IPv6len {
		return 0
	}

	return uint64(ip[15]) | uint64(ip[14])<<8 | uint64(ip[13])<<16 | uint64(ip[12])<<24 |
		uint64(ip[11])<<32 | uint64(ip[10])<<40 | uint64(ip[9])<<48 | uint64(ip[8])<<56
}
