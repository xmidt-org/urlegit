// SPDX-FileCopyrightText: 2023 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package urlegit

import (
	"fmt"
	"net"
	"strings"
)

// ForbidSubnet returns an Option that disallows the provided subnet addresses
// using only rules that do not result in a network call, unless a resolver is
// provided. If a resolver is provided, it will be used to resolve the hostname
// and check the returned IP addresses for loopback addresses.  If the subnet is
// invalid then the Option will return an error.
func ForbidSubnet(subnet string, resolver ...Resolver) Option {
	return forbidSubnetsOption("ForbidSubnet", []string{subnet}, resolver...)
}

func ForbidSubnets(subnets []string, resolver ...Resolver) Option {
	return forbidSubnetsOption("ForbidSubnets", subnets, resolver...)
}

func forbidSubnetsOption(name string, subnets []string, resolver ...Resolver) Option {
	f := forbidSubnetOption{
		optName:   "ForbidSubnet",
		subnets:   make([]*net.IPNet, 0, len(subnets)),
		originals: subnets,
	}

	for _, subnet := range subnets {
		_, cidr, err := net.ParseCIDR(subnet)
		if err != nil {
			return Error(fmt.Errorf("%w: invalid subnet '%s'", ErrInvalidInput, subnet))
		}
		f.subnets = append(f.subnets, cidr)
	}

	switch len(resolver) {
	case 0:
	case 1:
		f.r = resolver[0]
	default:
		return Error(fmt.Errorf("%w: only one resolver allowed", ErrInvalidInput))
	}

	return &f
}

type forbidSubnetOption struct {
	optName   string
	originals []string
	subnets   []*net.IPNet
	r         Resolver
}

func (n forbidSubnetOption) String() string {
	b := strings.Builder{}

	b.WriteString(n.optName)
	b.WriteString("(")
	comma := ""
	for _, original := range n.originals {

		b.WriteString(comma)
		b.WriteString("'")
		b.WriteString(original)
		b.WriteString("'")
		comma = ", "
	}
	if n.r != nil {
		b.WriteString(", resolver")
	}
	b.WriteString(")")

	return b.String()
}

func (n forbidSubnetOption) apply(c *Checker) {
	c.ipRules = append(c.ipRules, forbidSubnets(n.subnets))
	if n.r != nil {
		c.hostRules = append(c.hostRules, forbidSubnetsUser(n.subnets, n.r))
	}
}

func forbidSubnets(subnets []*net.IPNet) IPVador {
	return func(ip *net.IP) error {
		for _, subnet := range subnets {
			if subnet.Contains(*ip) {
				return ErrSubnetNotAllowed
			}
		}
		return nil
	}
}

func forbidSubnetsUser(subnets []*net.IPNet, fn Resolver) HostVador {
	return func(host string) error {
		ips, err := fn(host)
		if err != nil {
			return err
		}

		for _, ip := range ips {
			for _, subnet := range subnets {
				if subnet.Contains(ip) {
					return ErrSubnetNotAllowed
				}
			}
		}
		return nil
	}
}
