// SPDX-FileCopyrightText: 2023 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package urlegit

import (
	"fmt"
	"net"
)

// ForbidLoopback returns an Option that disallows loopback addresses using only
// rules that do not result in a network call, unless a resolver is provided.
// If a resolver is provided, it will be used to resolve the hostname and
// check the returned IP addresses for loopback addresses.
func ForbidLoopback(resolver ...Resolver) Option {
	switch len(resolver) {
	case 0:
		return forbidLoopbackOption(nil)
	case 1:
		return forbidLoopbackOption(resolver[0])
	default:
	}

	return Error(fmt.Errorf("%w: only one resolver allowed", ErrInvalidInput))
}

type forbidLoopbackOption Resolver

func (n forbidLoopbackOption) String() string {
	if n == nil {
		return "ForbidLoopback()"
	}
	return "ForbidLoopback(resolver)"
}

func (n forbidLoopbackOption) apply(c *Checker) {
	c.ipRules = append(c.ipRules, forbidLoopbackIP)
	c.hostRules = append(c.hostRules, forbidLoopbackHostname)
	if n != nil {
		c.hostRules = append(c.hostRules, forbidLoopbackUser(Resolver(n)))
	}
}

func forbidLoopbackIP(ip *net.IP) error {
	if ip.IsLoopback() {
		return ErrLoopback
	}
	return nil
}

func forbidLoopbackHostname(host string) error {
	if host == "localhost" {
		return ErrLoopback
	}
	return nil
}

func forbidLoopbackUser(fn Resolver) hostVador {
	return func(host string) error {
		ips, err := fn(host)
		if err != nil {
			return err
		}

		for _, ip := range ips {
			if ip.IsLoopback() {
				return ErrLoopback
			}
		}

		return nil
	}
}
