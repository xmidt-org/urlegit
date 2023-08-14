// SPDX-FileCopyrightText: 2023 Comcast Cable Communications Management, LLC
// SPDX-License-Identifier: Apache-2.0

package urlegit

import (
	"net"
)

// ForbidLoopback returns an Option that disallows loopback addresses using only
// rules that do not result in a network call.
func ForbidLoopback() Option {
	return &forbidLoopbackOption{}
}

type forbidLoopbackOption struct{}

func (forbidLoopbackOption) String() string {
	return "ForbidLoopback()"
}

func (forbidLoopbackOption) apply(c *Checker) {
	c.ipRules = append(c.ipRules, forbidLoopbackIP)
	c.hostRules = append(c.hostRules, forbidLoopbackHostname)
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
